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
   ./interpreter main.monkey
   ```

### Example

Create a file named `example.txt` with the following content:
   ```
   let x = 5;
   let y = 10;
   puts(x * y)
   ```

Run the interpreter:
   ```
   ./interpreter main.monkey
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thorsten Ball for the inspirational book *Writing an Interpreter in Go*.
- The Go Programming Language for the powerful language used.
- Me
- ChatGpt for this readme ( I truly cannot )
