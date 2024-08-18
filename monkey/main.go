package main

import (
	"BadlyWrittenStylesheet/Monkey/monkey/evaluator"
	"BadlyWrittenStylesheet/Monkey/monkey/lexer"
	"BadlyWrittenStylesheet/Monkey/monkey/object"
	"BadlyWrittenStylesheet/Monkey/monkey/parser"
	"fmt"
	"os"
)

func main() {
    filePath := os.Args[1]
    code, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("File %s does not exist or can not be opened?\n", filePath)
        return
    }
	l := lexer.New(string(code))
	p := parser.New(l)
	program := p.ParseProgram()
    env := object.NewEnviroment()
    res := evaluator.Eval(program, env)
    fmt.Println(res)
}

