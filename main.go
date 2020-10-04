package main

import (
	_"fmt"
	"os"
	"github.com/baybaraandrey/elephant/repl"
	"github.com/fatih/color"
)

const (
	elephant = `
___________.__                .__                   __   
\_   _____/|  |   ____ ______ |  |__ _____    _____/  |_ 
 |    __)_ |  | _/ __ \\____ \|  |  \\__  \  /    \   __\
 |        \|  |_\  ___/|  |_> >   Y  \/ __ \|   |  \  |  
/_______  /|____/\___  >   __/|___|  (____  /___|  /__|  
        \/           \/|__|        \/     \/     \/      
`
	version = "v0.0.1" 
)

func main() {
	color.Blue("%s%s", elephant, version)
	repl.Start(os.Stdin, os.Stdout)
}
