# Monkey Language Interpreter in Golang
## Overview

My Interpreter is a simple programming language interpreter written in Go, inspired by Thorsten Ball's *Writing an Interpreter in Go*. This project demonstrates fundamental concepts of language design, including lexical analysis, parsing, and interpretation.

## Features

- **Lexical Analysis**: Tokenizes source code into meaningful symbols.
- **Parsing**: Converts tokens into an Abstract Syntax Tree (AST).
- **Interpretation**: Executes the AST and evaluates expressions.
- **Interactive REPL**: Allows you to test the language interactively.

## Getting Started

### Prerequisites

- Go version something?

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/BadlyWrittenStylesheet/Monkey.git
   ```

2. Navigate to the project directory:
   ```
   cd Monkey
   ```

3. Build the project:
   ```
   go build -o monkey
   ```

### Usage

Run the interpreter with:
   ```
   ./monkey banana.monkey
   ```

### Example

Create a file named `banana.monkey` with the following content:
   ```
   let x = 5;
   let y = 10;
   if x > y {
      puts("Hello!")
   } else {
      puts("World!")
   }
   puts(2 + 3);
   let arr = [1, 2, 3];
   puts(arr[2]);
   let person = {"name": "Julian", "age": -1, "interests": ["cryptography", "math", "programming"]};
   let getFirstInterest = fn(x) {
       return x["interests"][0];
   };
   puts(getFirstInterest(person));
   puts(person["interests"][0]);
   ```

Run the interpreter:
   ```
   ./monkey banana.monkey
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thorsten Ball for the inspirational book *Writing an Interpreter in Go*. He is such a smart guy, really.
- The Go Programming Language for the powerful language used.
- Me
- ChatGPT on behalf of ClosedAI for this readme ( I truly cannot )
