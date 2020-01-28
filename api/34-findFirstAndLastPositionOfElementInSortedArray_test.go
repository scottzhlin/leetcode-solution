package api

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1, 2}, 2, []int{1, 1}},
	}

	for _, v := range tests {
		got := searchRange(v.nums, v.target)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
