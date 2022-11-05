package sink_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/micheam/go-stream/sink"
	"github.com/micheam/go-stream/source"
)

func TestDrain(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	src := source.FromSlice(ctx, want)

	if diff := cmp.Diff(want, sink.Collect(src)); diff != "" {
		t.Errorf("Drained mismatch (-want, +got):%s\n", diff)
	}
}
