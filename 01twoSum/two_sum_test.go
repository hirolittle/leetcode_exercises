package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"positive_nums", []int{1, 2, 5, 4}, 5, []int{0, 3}},
		{"negative_nums", []int{-3, -2, -2, 0}, -3, []int{0, 3}},
		{"mix_positive_negative", []int{-10, 4, 3, 10}, 0, []int{0, 3}},
		{"single_pair", []int{1, 2}, 3, []int{0, 1}},
		{"no_solution", []int{1, 2, 3}, 6, nil},
		{"multiple_pairs_first", []int{1, 4, 2, 3}, 4, []int{0, 3}},
		{"zeros", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"duplicates", []int{3, 3, 4, 5}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", got, tt.target, got, tt.want)
			}
		})
	}
}

// todo: benchmark test
