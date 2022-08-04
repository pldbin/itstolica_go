package server_test

import (
	"testing"

	"github.com/pldbin/prog1/internal/server"
)

func TestAdd(t *testing.T) {
	var (
		exp1 int
		// exp2 float64
	)

	sliceInt := []int{3, 5, 6, 7, 8}
	exp1 = 29
	got1, _ := server.Add(sliceInt...)
	if exp1 != got1 {
		t.Error("Expected ", exp1, ", got ", got1)
	}

}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v, _ := server.Add(pair.values...)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
