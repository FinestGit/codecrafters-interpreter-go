package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
	COMMA       rune = ','
	DOT         rune = '.'
	MINUS       rune = '-'
	PLUS        rune = '+'
	SEMICOLON   rune = ';'
	STAR        rune = '*'
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	errorsPresent := false

	for _, token := range string(fileContents) {
		if !errorsPresent {
			errorsPresent = scanToken(token)
		} else {
			scanToken(token)
		}
	}

	fmt.Println("EOF  null")
	if errorsPresent {
		os.Exit(65)
	}
}

func scanToken(token rune) bool {
	switch token {
	case LEFT_PAREN:
		fmt.Println("LEFT_PAREN ( null")
		return false
	case RIGHT_PAREN:
		fmt.Println("RIGHT_PAREN ) null")
		return false
	case LEFT_BRACE:
		fmt.Println("LEFT_BRACE { null")
		return false
	case RIGHT_BRACE:
		fmt.Println("RIGHT_BRACE } null")
		return false
	case COMMA:
		fmt.Println("COMMA , null")
		return false
	case DOT:
		fmt.Println("DOT . null")
		return false
	case MINUS:
		fmt.Println("MINUS - null")
		return false
	case PLUS:
		fmt.Println("PLUS + null")
		return false
	case SEMICOLON:
		fmt.Println("SEMICOLON ; null")
		return false
	case STAR:
		fmt.Println("STAR * null")
		return false
	default:
		fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %s\n", string(token))
		return true
	}
}
