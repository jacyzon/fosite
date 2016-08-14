package oauth2

import "time"

type Lifespan interface {
	GetAuthorizeCodeLifespan() time.Duration;
	GetAccessTokenLifespan() time.Duration;
	GetRefreshTokenLifespan() time.Duration;
}
