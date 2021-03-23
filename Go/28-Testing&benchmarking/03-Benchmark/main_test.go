package main

import "testing"

func TestCalc(t *testing.T)  {
	n := calc(1,5)
	if n != 6 {
		t.Error("Err")
	}
}

func BenchmarkCalc(b *testing.B)  {
	for i:=0; i <b.N; i++ {
		calc(56,5,1,5,96,51,58,9,1,9,18,9,8,19,18,9,8,19)
	}
}