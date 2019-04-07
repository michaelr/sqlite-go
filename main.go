package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("db > ")
		input := readInput()
		if strings.HasPrefix(input, ".") {
			if err := doMetaCmd(input); err != nil {
				fmt.Println(err)
			}
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

func doMetaCmd(cmd string) error {
	switch cmd {
	case ".exit":
		os.Exit(0)
	default:
		return errors.New("Unrecognized command")
	}
	return nil
}
