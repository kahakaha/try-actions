package main

import "fmt"

type Chirping interface {
	Chirp(msg string)
}

type Bird struct{}

func (bird Bird) Chirp(_ string) {
	fmt.Println("Hi")
}
