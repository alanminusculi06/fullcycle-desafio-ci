package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(1, 2)
	if res != 3 {
		t.Errorf("Sum(1, 2) = %d; expected 3", res)
	}
}
