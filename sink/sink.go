package sink

func Collect[T any](src <-chan T) []T {
	out := []T{}
	for v := range src {
		out = append(out, v)
	}
	return out
}
