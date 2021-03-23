package main

import "testing"

func TestMySum(t *testing.T)  {
	type test struct {
		data []int
		answer int
	}
	tests := []test{
		test{data: []int{1,5}, answer: 6},
		test{data: []int{-4,5}, answer: 1},
		test{data: []int{51,4}, answer: 55},
		test{data: []int{1,5}, answer: 6},
	}

	for _, v := range tests {
		s := mySum(v.data...)
		if s != v.answer {
			t.Error("Expected", v.answer, " :", s)
		}
	}
}