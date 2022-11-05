package testhelper

import (
	"context"
	"time"
)

const deadline = 200 * time.Millisecond

func WithDeadline(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithDeadline(context.Background(), time.Now().Add(deadline))
}
