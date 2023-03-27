package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}
	v := IsOne(i)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
