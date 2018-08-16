package main

import "testing"

func TestCI(t *testing.T) {
	a := 2
	b := 3
	c := a + b
	if c != 4 {
		panic("wrong c result")
	}
}

func TestCISecond(t *testing.T) {
	a := 2
	b := 3
	c := a * b
	if c != 6 {
		panic("wrong c result")
	}
}
