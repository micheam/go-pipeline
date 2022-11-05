package flow

import (
	"context"
	"math/rand"
)

func Broadcast[T any](ctx context.Context, n int, src <-chan T) []chan T {
	out := make([]chan T, n)
	for i := 0; i < n; i++ {
		out[i] = make(chan T)
	}
	go func() {
		defer func() {
			for i := range out {
				close(out[i])
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-src:
				if !ok {
					return
				}
				out[rand.Intn(n)] <- v
			}
		}
	}()
	return out
}
