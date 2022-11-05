package stream

import (
	"context"
	"testing"
	"time"
)

func TestTake(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	src := make(chan int)
	want := 100
	go func() {
		defer close(src)
		for i := 0; i < 1000; i++ {
			i := i
			src <- i
		}
	}()

	got := Collect(Take(ctx, want, src))
	if len(got) != want {
		t.Errorf("want %d but got %d", want, len(got))
	}
}
