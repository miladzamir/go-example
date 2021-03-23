package main

import "testing"

func TestMySum(t *testing.T)  {
	s := mySum(2,3)
	if s != 5 {
		t.Error("Expected", 5, " :", s)
	}
}