package sink

import "context"

func Collect[T any](ctx context.Context, src <-chan T) []T {
	out := []T{}
	for {
		select {
		case <-ctx.Done():
			return out
		case v, ok := <-src:
			if !ok {
				return out
			}
			out = append(out, v)
		}
	}
}
