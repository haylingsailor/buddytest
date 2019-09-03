package main

import "fmt"

// Speaker can say things.
type Speaker struct {
}

// SayHello says hello
func (s *Speaker) SayHello() string {
	return "hello-broken"
}

func main() {
	h := &Speaker{}
	fmt.Println(h.SayHello())
}
