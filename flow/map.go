package flow

import (
	"context"
	"fmt"
	"os"
)

func Map[I, O any](ctx context.Context, input <-chan I, fn func(context.Context, I) (O, error)) <-chan O {
	out := make(chan O)
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
				vv, err := fn(ctx, v)
				if err != nil {
					// TODO: design errorhandling
					fmt.Fprintf(os.Stderr, "[Map] %v\n", err)
					return
				}
				out <- vv
			}
		}
	}()
	return out
}
