package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/baybaraandrey/elephant/config"
	"github.com/baybaraandrey/elephant/evaluator"
	"github.com/baybaraandrey/elephant/lexer"
	"github.com/baybaraandrey/elephant/object"
	"github.com/baybaraandrey/elephant/parser"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer, conf config.Config) {
	env := object.NewEnvironment()

	if conf.Mode == config.INTERACTIVE {
		scanner := bufio.NewScanner(in)
		for {
			fmt.Printf(PROMPT)
			scanned := scanner.Scan()
			if !scanned {
				return
			}

			line := scanner.Text()

			l := lexer.New(line)
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParseErrors(out, p.Errors())
				continue
			}

			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	} else {
		input, err := ioutil.ReadAll(in)
		if err != nil {
			log.Fatal(err)
		}
		l := lexer.New(string(input))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
		}
		evaluator.Eval(program, env)
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
