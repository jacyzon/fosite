package oauth2

import (
	"strings"

	"github.com/ory-am/fosite"
	"github.com/ory-am/fosite/token/jwt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"time"
)

// RS256JWTStrategy is a JWT RS256 strategy.
type RS256JWTStrategy struct {
	*jwt.RS256JWTStrategy
}

func (h RS256JWTStrategy) signature(token string) string {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return ""
	}

	return split[2]
}

func (h RS256JWTStrategy) AccessTokenSignature(token string) string {
	return h.signature(token)
}

func (h RS256JWTStrategy) RefreshTokenSignature(token string) string {
	return h.signature(token)
}

func (h RS256JWTStrategy) AuthorizeCodeSignature(token string) string {
	return h.signature(token)
}

func (h *RS256JWTStrategy) GenerateAccessToken(_ context.Context, requester fosite.Requester) (token string, signature string, err error) {
	return h.generate(requester, "access_token")
}

func (h *RS256JWTStrategy) ValidateAccessToken(_ context.Context, _ fosite.Requester, token string) error {
	return h.validate(token)
}

func (h *RS256JWTStrategy) GenerateRefreshToken(_ context.Context, requester fosite.Requester) (token string, signature string, err error) {
	return h.generate(requester, "refresh_token")
}

func (h *RS256JWTStrategy) ValidateRefreshToken(_ context.Context, _ fosite.Requester, token string) error {
	return h.validate(token)
}

func (h *RS256JWTStrategy) GenerateAuthorizeCode(_ context.Context, requester fosite.Requester) (token string, signature string, err error) {
	return h.generate(requester, "authorization_code")
}

func (h *RS256JWTStrategy) ValidateAuthorizeCode(_ context.Context, requester fosite.Requester, token string) error {
	return h.validate(token)
}

func (h *RS256JWTStrategy) validate(token string) error {
	t, err := h.RS256JWTStrategy.Decode(token)
	if err != nil {
		return err
	}

	claims := jwt.JWTClaimsFromMap(t.Claims)
	if claims.IsNotYetValid() || claims.IsExpired() {
		return errors.New("Token claims did not validate")
	}

	return nil
}

func (h *RS256JWTStrategy) generate(requester fosite.Requester, tokenType string) (string, string, error) {
	if jwtSession, ok := requester.GetSession().(JWTSessionContainer); !ok {
		return "", "", errors.New("Session must be of type JWTSessionContainer")
	} else {
		claims := jwtSession.GetJWTClaims()
		if claims == nil {
			return "", "", errors.New("GetTokenClaims() must not be nil")
		} else {
			if session, ok := requester.GetSession().(fosite.Lifespan); !ok {
				return "", "", errors.New("Session must be of type Lifespan")
			} else {
				var duration time.Duration
				switch tokenType {
				case "access_token":
					duration = session.GetAccessTokenLifespan()
				case "refresh_token":
					duration = session.GetRefreshTokenLifespan()
				case "authorization_code":
					duration = session.GetAuthorizeCodeLifespan()
				}
				claims.ExpiresAt = time.Now().Add(duration)
				return h.RS256JWTStrategy.Generate(jwtSession.GetJWTClaims(), jwtSession.GetJWTHeader())
			}
		}
	}
}
