package weather

import "testing"

const (
	testV         = 32.1
	testC Celsius = testV
)

type Value interface {
	Unit() string
	Amount() float64
}

//go:noinline
func valueOf(c Celsius) float64 {
	return c.Amount()
}

//go:noinline
func valueOfInterface(v Value) float64 {
	return v.Amount()
}

func BenchmarkConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := valueOf(testC)
		if v != testV {
			b.Fatalf("unexpected value: got %f, want %f", v, testV)
		}
	}
}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := valueOfInterface(testC)
		if v != testV {
			b.Fatalf("unexpected value: got %f, want %f", v, testV)
		}
	}
}
