package flow_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/micheam/go-stream/flow"
	. "github.com/micheam/go-stream/internal/testinghelper"
	"github.com/micheam/go-stream/sink"
	"github.com/micheam/go-stream/source"
)

func TestMap(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()

	values := []string{"x", "xxx", "xxxxx"}
	want := []int{1, 3, 5}

	fn := func(_ context.Context, v string) (int, error) {
		return len(v), nil
	}

	src := source.FromSlice(ctx, values)
	mapped := flow.Map(ctx, src, fn)
	got := sink.Collect(ctx, mapped)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mapped mismatch (-want, +got):%s\n", diff)
	}
}

func TestMap_cancel(t *testing.T) {
	ctx, cancel := WithDeadline(context.Background())
	defer cancel()

	values := []string{"x", "xxx", "xxxxx"}
	want := []int{1, 3}

	fn := func(_ context.Context, v string) (int, error) {
		if v == "xxxxx" {
			return 0, errors.New("foo")
		}
		return len(v), nil
	}

	src := source.FromSlice(ctx, values)
	mapped := flow.Map(ctx, src, fn)
	got := sink.Collect(ctx, mapped)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mapped mismatch (-want, +got):%s\n", diff)
	}
}
