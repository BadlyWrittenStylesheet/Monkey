package main

import (
	"fmt"
	"os"
	"os/user"
	"BadlyWrittenStylesheet/Monkey/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the ultimate programming language, MONKEY\n", user.Username)
	fmt.Printf("Feel free to write some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

