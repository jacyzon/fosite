package compose

import (
)

type Config struct {
	*Lifespan
	// HashCost sets the cost of the password hashing cost. Defaults to 12.
	HashCost int
}

// GetHashCost returns how many iterations the bcrypt will use.
func (c *Config) GetHashCost() int {
	if c.HashCost == 0 {
		return 12
	}
	return c.HashCost
}
