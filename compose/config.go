package compose

import "time"

type Config struct {
	// AccessTokenLifespan sets how long an access token is going to be valid. Defaults to one hour.
	AccessTokenLifespan   time.Duration

	// RefreshTokenLifespan sets how long an refresh token is going to be valid. Defaults to one week.
	RefreshTokenLifespan  time.Duration

	// AuthorizeCodeLifespan sets how long an authorize code is going to be valid. Defaults to fifteen minutes.
	AuthorizeCodeLifespan time.Duration

	// IDTokenLifespan sets how long an id token is going to be valid. Defaults to one hour.
	IDTokenLifespan       time.Duration

	// HashCost sets the cost of the password hashing cost. Defaults to 12.
	HashCost              int
}

// GetAuthorizeCodeLifespan returns how long an authorize code should be valid. Defaults to one fifteen minutes.
func (c *Config) GetAuthorizeCodeLifespan() time.Duration {
	if c.AuthorizeCodeLifespan == 0 {
		return time.Minute * 15
	}
	return c.AuthorizeCodeLifespan
}

// GetAccessTokenLifespan returns how long a access token should be valid. Defaults to one hour.
func (c *Config) GetAccessTokenLifespan() time.Duration {
	if c.AccessTokenLifespan == 0 {
		return time.Hour
	}
	return c.AccessTokenLifespan
}

// GetRefreshTokenLifespan returns how long an refresh code should be valid. Defaults to two hour.
func (c *Config) GetRefreshTokenLifespan() time.Duration {
	if c.RefreshTokenLifespan == 0 {
		return time.Hour * 24 * 7
	}
	return c.RefreshTokenLifespan
}

// GeIDTokenLifespan returns how long an id token should be valid. Defaults to one hour.
func (c *Config) GetIDTokenLifespan() time.Duration {
	if c.IDTokenLifespan == 0 {
		return time.Hour
	}
	return c.IDTokenLifespan
}

// GetHashCost returns how many iterations the bcrypt will use.
func (c *Config) GetHashCost() int {
	if c.HashCost == 0 {
		return 12
	}
	return c.HashCost
}
