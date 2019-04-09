package fibo_test

import (
	"reflect"
	"testing"

	"github.com/davyj0nes/euler-problems-golang/2-even-fibonacci-numbers/fibo"
)

func TestGetFibonacci(t *testing.T) {
	tests := []struct {
		name string
		max  int
		want int
	}{
		{
			name: "max of 1",
			max:  1,
			want: 0,
		},
		{
			name: "max of 2",
			max:  2,
			want: 2,
		},
		{
			name: "max of 90",
			max:  90,
			want: 44,
		},
		{
			name: "max of 4 000 000",
			max:  4000000,
			want: 4613732,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibo.SumEven(tt.max)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: (%d), got: (%d)", tt.want, got)
			}
		})
	}
}
