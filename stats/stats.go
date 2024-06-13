package stats

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

// Max - returns the maximum value from the slice
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice provides")
	}

	max := values[0]

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max, nil
}
