package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/baybaraandrey/elephant/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Elephant programming language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
