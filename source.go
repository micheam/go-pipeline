package stream

import "context"

// Stream is a stream generator, which converts the received values into a stream and returns it.
func Stream[T any](ctx context.Context, values []T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()
	return out
}

func StreamFn[T any](ctx context.Context, fn func() T) <-chan T {
	dest := make(chan T)
	go func() {
		defer close(dest)
		for {
			select {
			case <-ctx.Done():
				return
			case dest <- fn():
			}
		}
	}()
	return dest
}
