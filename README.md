# Monkey Language ( currently only lexer and parser )
## How to "experience it" ( requires the go compiler and stuff )
1.  Clone the repo
2.  Go inside the direcory `cd ./Monkey/monkey`
3.  Run `go run main.go`
4.  You are prompted with a ?>
5.  Try typing 1 + 2 * 3 / 4 + 5 / 6 <enter>
6.  It will print this expression with parentheses to tell what is executed first
7.  You can try function calls and operators too, like `1 + add(2 * 3, 4) / 3 > 8`
8.  And we get this: ((1 + (add((2 * 3), 4) / 3)) > 8) see the parentheses around 2 * 3, it works!
9.  Check it out and bye!
10.  Shhh... you can try `let foo = bar != 5 * 2;` or `return foo > 5 + bar;` but it sometimes breaks? Remember semicolons ;)
11.  Oh and errors, `let = true` will give u `expectred next token to be IDENT. got = instead` 😎
12.  Now, bye!
