package source

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

func FromReader(ctx context.Context, input io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		scan := bufio.NewScanner(input)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scan.Scan() {
					out <- scan.Text()
					continue
				}
				if err := scan.Err(); err != nil {
					// TODO: design errorhandling
					fmt.Fprintf(os.Stderr, "[source] FromReader: %v", err)
					return
				}
			}
		}
	}()
	return out
}
