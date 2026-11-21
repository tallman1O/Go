package main

import (
	"fmt"
	"time"
)

// say prints a string 5 times, pausing between each print.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // pause so the two loops can interleave
		fmt.Println(s)
	}
}

func main() {
	// go starts a new goroutine — a lightweight concurrent function call.
	// "world" runs in the background while main continues.
	go say("world")

	// This runs in the main goroutine (no "go" keyword).
	// Without this call, main would exit immediately and kill the "world" goroutine.
	say("hello")

	// When main returns, the program ends — including any leftover goroutines.
}
