package main

import (
	"bufio"
	"os"

	"github.com/fadilix/couik/internal/game"
	"github.com/fadilix/couik/pkg/typing"
	"golang.org/x/term"
)

func main() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	target := typing.GetRandomQuote()
	reader := bufio.NewReader(os.Stdin)

	game.GameLoop(target, reader)
}
