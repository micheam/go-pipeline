package pipeline

import (
	"context"
	"math/rand"
)

// TODO: Fan-out
// - [x] Broadcast[T] – (1 input, N outputs) given an input element emits to each output
// - [ ] Balance[T] – (1 input, N outputs) given an input element emits to one of its output ports
// - [ ] UnzipWith[In,A,B,...] – (1 input, N outputs) takes a function of 1 input that given a value for each input emits N output elements (where N <= 20)
// - [ ] UnZip[A,B] – (1 input, 2 outputs) splits a stream of (A,B) tuples into two streams, one of type A and one of type B

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
