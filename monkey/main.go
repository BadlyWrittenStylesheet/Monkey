package main

import (
	"fmt"
	"os"
	"os/user"
	"BadlyWrittenStylesheet/Monkey/monkey/repl"
	"BadlyWrittenStylesheet/Monkey/monkey/lexer"
	"BadlyWrittenStylesheet/Monkey/monkey/token"
)

func main() {
	tests := []struct {
		input	string
		expected string
	}{
		{
	"-a * b",
	"((-a) * b)",
},
{
	"!-a",
	"(!(-a))",
},
{
	"a + b + c",
	"((a + b) + c)",
},
{
	"a + b - c",
	"((a + b) - c)",
},
{
	"a * b * c",
	"((a * b) * c)",
},
{
	"a * b / c",
	"((a * b) / c)",
},
{
	"a + b / c",
	"(a + (b / c))",
},
{
	"a + b * c + d / e - f",
	"(((a + (b * c)) + (d / e)) - f)",
},
{
	"3 + 4; -5 * 5",
	"(3 + 4)((-5) * 5)",
},
{
	"5 > 4 == 3 < 4",
	"((5 > 4) == (3 < 4))",
},
{
	"5 < 4 != 3 > 4",
	"((5 < 4) != (3 > 4))",
},
{
	"((5 < 4) |= (3 > 4))",
	"3 + 4 * 5 == 3 * 1 + 4 * 5",
},
{
	"3 + 4 * 5 == 3 * 1 + 4 * 5",
	"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
},
}

	for _, t := range tests {
		l := lexer.New(t.expected)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
	fmt.Print(tests)
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the ultimate programming language, MONKEY\n", user.Username)
	fmt.Printf("Feel free to write some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

