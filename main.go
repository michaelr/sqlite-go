package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("db > ")
		input := readInput()

		switch input {
		case ".exit":
			os.Exit(0)
		default:
			fmt.Printf("Unrecognized command '%s'.\n", input)
		}
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Unable to read input\n")
	}

	return strings.TrimSpace(text)
}
