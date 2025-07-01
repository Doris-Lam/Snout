package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMPT = ">> "

const MONKEY_FACE = `
  / \\__
 (    @\\___
 /         O
/   (_____/
/_____/ U
`

const DOG_FACE = `
  / \\__
 (    @\\___
 /         O
/   (_____/
/_____/ U
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, DOG_FACE)
	io.WriteString(out, "Oups ! Nous avons rencontré un problème avec Snout !\n")
	io.WriteString(out, "Erreurs d'analyse :\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			// Print DOG_FACE if it's an error
			if _, ok := evaluated.(*object.Error); ok {
				io.WriteString(out, DOG_FACE)
			}
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
