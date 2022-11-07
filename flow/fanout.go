package flow

import (
	"context"
)

func Broadcast[T any](ctx context.Context, n int, src <-chan T) []<-chan T {
	consumer := func(ctx context.Context, in <-chan T) <-chan T {
		_out := make(chan T)
		go func() {
			defer close(_out)
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					_out <- v
				}
			}
		}()
		return _out
	}

	out := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		out[i] = consumer(ctx, src)
	}
	return out
}
