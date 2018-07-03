package main

import "testing"

func TestAdd(t *testing.T) {
	var c = add(10, 20)
	if c != 31 {
		t.Error("期待値と一致しない")
	}
}
