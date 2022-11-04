package pipeline

import (
	"context"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	var (
		values      = 10
		channels    = 3
		src         = make([]chan int, channels)
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	)
	defer cancel()

	for i := 0; i < channels; i++ {
		src[i] = make(chan int)
	}

	go func() {
		for i := 0; i < channels; i++ {
			for j := 0; j < values; j++ {
				src[i] <- j
			}
		}
		for i := 0; i < channels; i++ {
			close(src[i])
		}
	}()

	coll := Collect(Merge(ctx, src))
	if len(coll) != values*channels {
		t.Errorf("want %d, but got %d", values*channels, len(coll))
	}
}
