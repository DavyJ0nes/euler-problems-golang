package multiples_test

import (
	"testing"

	"github.com/davyj0nes/euler-problems-golang/1-multiples-of-3-and-5/multiples"
)

func TestMultiplesOf3Or5(t *testing.T) {
	tests := []struct {
		name string
		max  int
		want int
	}{
		{
			name: "max of 10",
			max:  10,
			want: 23,
		},
		{
			name: "max of 100",
			max:  100,
			want: 2318,
		},
		{
			name: "max of 1000",
			max:  1000,
			want: 233168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := multiples.NewOf3Or5(tt.max)
			got := m.Sum()

			if got != tt.want {
				t.Fatalf("want: (%d), got: (%d)", tt.want, got)
			}
		})
	}
}
