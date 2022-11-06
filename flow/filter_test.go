package flow_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/micheam/go-stream/flow"
	. "github.com/micheam/go-stream/internal/testinghelper"
	"github.com/micheam/go-stream/sink"
	"github.com/micheam/go-stream/source"
)

func TestFilter(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()

	values := []int{1, 2, 3, 4, 5}
	want := []int{1, 3, 5}

	fn := func(_ context.Context, v int) (bool, error) {
		return v%2 != 0, nil
	}

	src := source.FromSlice(ctx, values)
	mapped := flow.Filter(ctx, src, fn)
	got := sink.Collect(ctx, mapped)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mapped mismatch (-want, +got):%s\n", diff)
	}
}

func TestFilter_Abortion(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()

	values := []int{1, 2, 3, 4, 5}
	want := []int{1, 3}

	fn := func(_ context.Context, v int) (bool, error) {
		if v == 4 {
			return false, flow.ErrAbort
		}
		return v%2 != 0, nil
	}

	src := source.FromSlice(ctx, values)
	mapped := flow.Filter(ctx, src, fn)
	got := sink.Collect(ctx, mapped)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mapped mismatch (-want, +got):%s\n", diff)
	}
}
