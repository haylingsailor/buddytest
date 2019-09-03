package main

import "testing"

func TestHello(t *testing.T) {
	sut := &Speaker{}

	r := sut.SayHello()
	if r != "hello" {
		t.Error("Didn't say hello")
	}
}
