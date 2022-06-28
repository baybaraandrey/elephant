package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/baybaraandrey/elephant/repl"
)

func main() {
	flag.Parse()
	args := flag.Args()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(args) > 0 {
		filePath := args[0]
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("%s", err)
		}

		repl.Start(file, os.Stdout, repl.NON_INTERACTIVE)
	} else {
		fmt.Printf("Hello %s! This is Elephant programming language.\n", user.Username)
		repl.Start(os.Stdin, os.Stdout, repl.INTERACTIVE)
	}
}
