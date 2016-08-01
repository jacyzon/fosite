package fosite

import (
	"time"
	"golang.org/x/net/context"
)

type Lifespan interface {
	GetLifespan(ctx context.Context, requester Requester, tokenType string) time.Duration
}

