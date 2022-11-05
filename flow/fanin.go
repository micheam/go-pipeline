package flow

import (
	"context"
	"sync"
)

func Merge[T any](ctx context.Context, src []chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	for _, ch := range src {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}(ch)
	}

	go func() { // Closer
		wg.Wait()
		close(out)
	}()

	return out
}
