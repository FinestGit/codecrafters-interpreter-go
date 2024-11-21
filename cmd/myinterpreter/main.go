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
	EQUAL       rune = '='
	GREATER     rune = '>'
	LESS        rune = '<'
	BANG        rune = '!'
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

	tokens := []rune(string(fileContents))

	currentPointer := 0
	nextPointer := 1

	for range len(tokens) {
		isEnd := nextPointer >= len(tokens)
		if nextPointer > len(tokens) {
			break
		}
		token := tokens[currentPointer]
		switch token {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case COMMA:
			fmt.Println("COMMA , null")
		case DOT:
			fmt.Println("DOT . null")
		case MINUS:
			fmt.Println("MINUS - null")
		case PLUS:
			fmt.Println("PLUS + null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		case STAR:
			fmt.Println("STAR * null")
		case EQUAL:
			if isEnd {
				fmt.Println("EQUAL = null")
			} else {
				if match(tokens[nextPointer]) {
					fmt.Println("EQUAL_EQUAL == null")
					nextPointer++
					currentPointer++
				} else {
					fmt.Println("EQUAL = null")
				}
			}
		case BANG:
			if isEnd {
				fmt.Println("BANG ! null")
			} else {
				if match(tokens[nextPointer]) {
					fmt.Println("BANG_EQUAL != null")
					nextPointer++
					currentPointer++
				} else {
					fmt.Println("BANG ! null")
				}
			}
		case GREATER:
			if isEnd {
				fmt.Println("GREATER > null")
			} else {
				if match(tokens[nextPointer]) {
					fmt.Println("GREATER_EQUAL >= null")
					nextPointer++
					currentPointer++
				} else {
					fmt.Println("GREATER > null")
				}
			}
		case LESS:
			if isEnd {
				fmt.Println("LESS < null")
			} else {
				if match(tokens[nextPointer]) {
					fmt.Println("LESS_EQUAL <= null")
					nextPointer++
					currentPointer++
				} else {
					fmt.Println("LESS < null")
				}
			}
		default:
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %s\n", string(token))
			errorsPresent = true
		}
		nextPointer++
		currentPointer++
	}

	fmt.Println("EOF  null")
	if errorsPresent {
		os.Exit(65)
	}
}

func match(token rune) bool {
	if token == EQUAL {
		return true
	} else {
		return false
	}
}
