package main

import (
	"fmt"
	"github.com/baybaraandrey/elephant/repl"
	"github.com/fatih/color"
	"os"
)

const (
	elephant = `___________.__                .__                   __   
\_   _____/|  |   ____ ______ |  |__ _____    _____/  |_ 
 |    __)_ |  | _/ __ \\____ \|  |  \\__  \  /    \   __\
 |        \|  |_\  ___/|  |_> >   Y  \/ __ \|   |  \  |  
/_______  /|____/\___  >   __/|___|  (____  /___|  /__|  
        \/           \/|__|        \/     \/     \/      
`
	version = "v0.0.1"
)

func main() {
	welcome()
	repl.Start(os.Stdin, os.Stdout)
}

func welcome() {
	color.Blue("%s%s", elephant, version)
	fmt.Println()
}
