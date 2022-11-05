package flow

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Filter[T any](ctx context.Context, input <-chan T, fn func(context.Context, T) (bool, error)) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-input:
				if !ok {
					return
				}
				ok, err := fn(ctx, v)
				if err != nil {
					if !errors.Is(err, ErrAbort) {
						fmt.Fprintf(os.Stderr, "[Filter] %v\n", err)
					}
					return
				}
				if !ok {
					continue
				}
				out <- v
			}
		}
	}()
	return out
}
