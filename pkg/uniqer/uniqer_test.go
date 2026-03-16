package uniquer

import (
	"reflect"
	"testing"
)

type TestCase struct {
	give []int
	want []int
}

func TestUniquer(t *testing.T) {

	cases := []TestCase{
		{[]int{1, 2, 3, 4, 4, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 3, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 5, 6}, []int{1, 3, 5, 6}},
	}

	for _, c := range cases {
		unique := Unique(c.give)
		if !reflect.DeepEqual(unique, c.want) {
			t.Errorf("uniqer got %v want %v", unique, c.want)
		}
	}
}

func TestUniquerOptimized(t *testing.T) {

	cases := []TestCase{
		{[]int{1, 2, 3, 4, 4, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 3, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 5, 6}, []int{1, 3, 5, 6}},
	}

	for _, c := range cases {
		unique := UniqueOptimized(c.give)
		if !reflect.DeepEqual(unique, c.want) {
			t.Errorf("uniqer got %v want %v", unique, c.want)
		}
	}
}
