package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const STDIN_PROMPT string = ">> "
const ERR_KEY_NOT_EXIST error = fmt.Errorf("Key does not exist")

type Store struct {
	Kv     map[string]string
	Counts map[string]int
	Parent *db
}

func NewStore() *Store {
	s := new(Store)
	s.KV = make(map[string]string)
	s.Counts = make(map[string]int)

	return s
}

func (s *Store) Set(key, value string) error {
	if valueOld, exists := s.Kv[key]; exists {
		s.Counts[valueOld] -= 1
	}

	s.Kv[key] = value
	s.Counts[key] += 1

	return nil
}

func (s *Store) Get(key string) (string, error) {
	value, exists := s.Kv[key]

	// Base Condition
	if exists {
		return value, nil
	}
	if s.Parent == nil {
		return value, ERR_KEY_NOT_EXIST
	}
	return s.Parent.Get(key)
}

func (s *Store) Delete(key string) error {
	value, exists := s.Kv[key]
	if !exists {
		return ERR_KEY_NOT_EXIST
	}
	kv.Counts[value] -= 1
	delete(s.Kv, key)
}

func (s *Store) Count(key string) (int, error) {
	return s.Counts[key], nil
}

var primaryStore *Store

func init() {
	primaryStore = NewStore()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(STDIN_PROMPT)

	for scanner.Scan() {
		stmt := scanner.Text()

		cmd, args := splitStatement(stmt)

		fn, exists := funcMap[cmd]
		if !exists {
			handleErrf("Invalid command: %s", cmd)
			continue
		}

		fn(args...)
	}
}

var funcMap map[string]func(...string) = map[string]func(...string){
	"SET": handleSet,
	"GET": handleGet,
}

func handleErrf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
	fmt.Print(STDIN_PROMPT)
}
func handleReturn(msg string, args ...interface{}) {
	if msg != "" {
		fmt.Printf(msg+"\n", args...)
	}
	fmt.Print(STDIN_PROMPT)
}

func handleSet(args ...string) {
	// Check the number of args
	var numArgsExpected int = 2
	if len(args) != numArgsExpected {
		handleErrf("Invalid number of arguments provided: expected %d got %d", numArgsExpected, len(args))
	}

	key := args[0]
	value := args[1]

	// get current value
	if _value, exists := kv[key]; exists {
		cnts[_value] = cnts[_value] - 1
	}

	kv[key] = value
	cnts[key] += 1

	handleReturn("")
}

func handleGet(args ...string) {
	// Check the number of args
	var numArgsExpected int = 1
	if len(args) != numArgsExpected {
		handleErrf("Invalid number of arguments provided: expected %d got %d", numArgsExpected, len(args))
	}

	key := args[0]
	value, exists := kv[key]
	if !exists {
		handleReturn("NULL")
	}

	handleReturn(value)
}

func handleDelete(args ...string) {
	// Check the number of args
	var numArgsExpected int = 1
	if len(args) != numArgsExpected {
		handleErrf("Invalid number of arguments provided: expected %d got %d", numArgsExpected, len(args))
	}
	key := args[0]

	delete(kv, key)
}

func validateStatement(q string) error {
	if strings.TrimSpace(q) == "" {
		return fmt.Errorf("Empty statement passed")
	}
	return nil
}

func splitStatement(stmt string) (string, []string) {
	// Split the statement by single whitespace
	stmtSplit := strings.Split(stmt, " ")
	// First word is the command, while the remaining (if any) are the args
	return strings.ToUpper(stmtSplit[0]), stmtSplit[1:]
}
