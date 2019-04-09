// sqlite-go is a go implementation of sqlite
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// rawSQL is a string of SQL
type rawSQL string

// RawSQL returns a string represention of the rawSQL
func (s rawSQL) RawSQL() string {
	return string(s)
}

// sqlStatement is sql statement that's been parsed and is ready to execute.
type sqlStatement interface {
	RawSQL() string
}
type insertStatement struct {
	rawSQL
}
type selectStatement struct {
	rawSQL
}

func main() {

	// This is the REPL loop
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
				if err := executeStatement(stmt); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// readInput reads from Stdin and trims the surrounding whitespace.
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic("Unable to read input\n")
	}

	return strings.TrimSpace(text)
}

// doMetaCmd executes sqlite commands that start with a "."
// Meta commands are non-sql. For example, ".exit"
func doMetaCmd(cmd string) error {
	switch cmd {
	case ".exit":
		os.Exit(0)
	default:
		return errors.New("Unrecognized command")
	}
	return nil
}

// prepareStatement parses a SQL string and returns the corresponding
// sqlStatement according to what sql statement it is. insert, select, etc.
func prepareStatement(input string) (sqlStatement, error) {
	switch {
	case strings.HasPrefix(input, "insert"):
		return &insertStatement{rawSQL(input)}, nil
	case strings.HasPrefix(input, "select"):
		return &selectStatement{rawSQL(input)}, nil
	default:
		return nil, fmt.Errorf("Unrecognized statement: %v", input)
	}
}

// executeStatement runs a sqlStatement
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
