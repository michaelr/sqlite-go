package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type sqlString string

func (s sqlString) rawSQL() string {
	return string(s)
}

type sqlStatement interface {
	rawSQL() string
}

type insertStatement struct {
	sqlString
}
type selectStatement struct {
	sqlString
}

func main() {
	for {
		fmt.Print("db > ")
		input := readInput()

		if input == "" {
			continue
		}
		if strings.HasPrefix(input, ".") {
			if err := doMetaCmd(input); err != nil {
				fmt.Println(err)
			}
		} else {
			if stmt, err := prepareStatement(input); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(stmt.rawSQL())
				if err := executeStatement(stmt); err != nil {
					fmt.Println(err)
				}
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

func prepareStatement(input string) (sqlStatement, error) {
	switch {
	case strings.HasPrefix(input, "insert"):
		return &insertStatement{sqlString(input)}, nil
	case strings.HasPrefix(input, "select"):
		return &selectStatement{sqlString(input)}, nil
	default:
		return nil, fmt.Errorf("Unrecognized statement: %v", input)
	}
}

func executeStatement(statement sqlStatement) error {
	switch statement.(type) {
	case *insertStatement:
		fmt.Println("TODO: execute insert")
	case *selectStatement:
		fmt.Println("TODO: execute select")
	default:
		return errors.New("Unknown statement type")
	}
	return nil
}
