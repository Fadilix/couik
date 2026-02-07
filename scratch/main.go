package main

import "fmt"

type typingMode int

const (
	quoteMode typingMode = iota
	timedMode
)

func main() {
	fmt.Println(typingMode(0))
	// fmt.Println(quoteMode)
}
