package fosite

import "time"

type Lifespan interface {
	GetAuthorizeCodeLifespan() time.Duration
	GetAccessTokenLifespan() time.Duration
	GetRefreshTokenLifespan() time.Duration
	GetIDTokenLifespan() time.Duration
}
