package stream

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDrain(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	src := Stream(ctx, want)

	if diff := cmp.Diff(want, Collect(src)); diff != "" {
		t.Errorf("Drained mismatch (-want, +got):%s\n", diff)
	}
}
