package pipeline

import (
	"context"
	"sync"
)

// Fan-in
//
// - [x] Merge[In] – (N inputs , 1 output) picks randomly from inputs pushing them one by one to its output
// - [ ] MergePreferred[In] – like Merge but if elements are available on preferred port, it picks from it, otherwise randomly from others
// - [ ] MergePrioritized[In] – like Merge but if elements are available on all input ports, it picks from them randomly based on their priority
// - [ ] MergeLatest[In] – (N inputs, 1 output) emits List[In], when i-th input stream emits element, then i-th element in emitted list is updated
// - [ ] MergeSequence[In] – (N inputs, 1 output) emits List[In], where the input streams must represent a partitioned sequence that must be merged back together in order
// - [ ] ZipWith[A,B,...,Out] – (N inputs, 1 output) which takes a function of N inputs that given a value for each input emits 1 output element
// - [ ] Zip[A,B] – (2 inputs, 1 output) is a ZipWith specialised to zipping input streams of A and B into a (A,B) tuple stream
// - [ ] Concat[A] – (2 inputs, 1 output) concatenates two streams (first consume one, then the second one)

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
