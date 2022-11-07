package flow_test

import (
	"context"
	"testing"
	"time"

	"github.com/micheam/go-stream/flow"
	"github.com/micheam/go-stream/sink"
)

func TestMerge(t *testing.T) {
	var (
		values      = 10
		channels    = 3
		src         = make([]<-chan int, channels)
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	)
	defer cancel()

	makech := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < values; i++ {
				out <- i
			}
			return
		}()
		return out
	}
	for i := 0; i < channels; i++ {
		src[i] = makech()
	}

	coll := sink.Collect(ctx, flow.Merge(ctx, src...))
	if len(coll) != values*channels {
		t.Errorf("want %d, but got %d", values*channels, len(coll))
	}
}
