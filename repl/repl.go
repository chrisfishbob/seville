package repl

import (
	"bufio"
	"fmt"
	"io"
	"seville/ast"
	"seville/compiler"
	"seville/evaluator"
	"seville/lexer"
	"seville/object"
	"seville/parser"
	"seville/vm"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, isCompiled bool) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	if isCompiled {
		fmt.Println("Executing using the experimental compiler ...")
	}

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "quit" {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if isCompiled {
			output, err := executeProgramWithCompiler(program)
			if err != nil {
				io.WriteString(out, err.Error())
				continue
			}
			io.WriteString(out, output)
		} else {
			output := executeProgramWithInterpreter(program, env)
			io.WriteString(out, output)
		}
	}
}

func executeProgramWithCompiler(program *ast.Program) (string, error) {
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		return "", fmt.Errorf("compilation failed:\n %s", err)
	}

	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		return "", fmt.Errorf("executing bytecode failed: \n %s", err)
	}

	stackTop := machine.StackTop()
	return stackTop.Inspect() + "\n", nil
}

func executeProgramWithInterpreter(program *ast.Program, env *object.Environment) string {
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		return evaluated.Inspect() + "\n"
	}
	return ""
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Oh no my program!\n")
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
