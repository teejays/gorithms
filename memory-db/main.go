package main

import (
	"bufio"
	"fmt"
	"github.com/teejays/clog"
	"os"
	"strings"
)

const STDIN_PROMPT string = ">> "

var ERR_INVALID_COMMAND_KEYWORD error = fmt.Errorf("INVALID COMMAND")
var ERR_INVALID_ARGS_NUM error = fmt.Errorf("INVALID NUMBER OF ARGUMENTS PROVIDED")
var ERR_STMT_EMPTY error = fmt.Errorf("EMPTY STATEMENT PROVIDED")

var EnableTestMode bool = false

var primaryStore *Store
var currentStore *Store

func init() {
	primaryStore = NewStore(nil)
	currentStore = primaryStore
}

type ActionHandler struct {
	Fn      func(...string) (string, error)
	NumArgs int
}

var funcMap map[string]ActionHandler = map[string]ActionHandler{
	"SET":      ActionHandler{handleSet, 2},
	"GET":      ActionHandler{handleGet, 1},
	"DELETE":   ActionHandler{handleDelete, 1},
	"COUNT":    ActionHandler{handleCount, 1},
	"BEGIN":    ActionHandler{handleBegin, 0},
	"ROLLBACK": ActionHandler{handleRollback, 0},
	"COMMIT":   ActionHandler{handleCommit, 0},
	"END":      ActionHandler{handleEnd, 0},
}

func main() {
	clog.LogLevel = 5

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(STDIN_PROMPT)

	for scanner.Scan() {
		stmt := scanner.Text()
		err := validateStatement(stmt)
		if err == ERR_STMT_EMPTY {
			fmt.Print(STDIN_PROMPT)
			continue
		}
		if err != nil {
			fmt.Println(err)
		}
		out, err := handleStatement(stmt)
		if err != nil {
			fmt.Println(err)
		}
		if out != "" {
			fmt.Println(out)
		}
		fmt.Print(STDIN_PROMPT)
	}
}

func handleStatement(stmt string) (string, error) {
	//clog.Debugf("[handleStatement] Statement: %s", stmt)
	cmd, args := splitStatement(stmt)
	h, exists := funcMap[cmd]
	if !exists {
		return "", ERR_INVALID_COMMAND_KEYWORD
	}

	// Check the number of args
	if len(args) != h.NumArgs {
		return "", ERR_INVALID_ARGS_NUM
	}

	out, err := h.Fn(args...)
	if err != nil {
		return out, err
	}

	clog.Debugf("CurrentStore: %v", currentStore)
	return out, nil
}

func handleSet(args ...string) (string, error) {
	key := args[0]
	value := args[1]

	err := currentStore.Set(key, value)
	return "", err
}

func handleGet(args ...string) (string, error) {
	key := args[0]

	val, err := currentStore.Get(key)
	if err == ERR_KEY_NOT_EXIST || val == "" {
		return "NULL", nil
	}
	return val, err
}

func handleDelete(args ...string) (string, error) {
	key := args[0]

	err := currentStore.Delete(key)
	return "", err
}

func handleCount(args ...string) (string, error) {
	val := args[0]

	cnt, err := currentStore.Count(val)
	return fmt.Sprintf("%d", cnt), err
}

func handleBegin(args ...string) (string, error) {
	// Whenever a transaction begins, we should create a new empty Store (not a copy)
	// which can store intermediate states of those things that have changed.
	trxnStore := NewStore(currentStore)
	currentStore = trxnStore

	return "", nil
}

func handleRollback(args ...string) (string, error) {
	if currentStore.Parent == nil {
		return "", fmt.Errorf("No transaction to rollback")
	}
	currentStore = currentStore.Parent
	return "", nil
}

func handleCommit(args ...string) (string, error) {
	if currentStore.Parent == nil {
		return "", fmt.Errorf("No transaction to commit")
	}

	// we need to merge the current store to the primary store
	for k, v := range currentStore.Kv {
		primaryStore.Kv[k] = v
	}
	for _, v := range primaryStore.Kv {
		cnt, err := currentStore.Count(v)
		if err != nil {
			return "", fmt.Errorf("Error while commiting transaction: %s", err)
		}
		primaryStore.CountDiff[v] = cnt
	}
	currentStore = currentStore.Parent
	return "", nil
}

func handleEnd(args ...string) (string, error) {
	if !EnableTestMode {
		os.Exit(0)
	} else {
		primaryStore = NewStore(nil)
		currentStore = primaryStore
	}
	return "", nil
}

func validateStatement(q string) error {
	if strings.TrimSpace(q) == "" {
		return ERR_STMT_EMPTY
	}
	return nil
}

func splitStatement(stmt string) (string, []string) {
	// Split the statement by single whitespace
	stmtSplit := strings.Split(stmt, " ")
	// First word is the command, while the remaining (if any) are the args
	return strings.ToUpper(stmtSplit[0]), stmtSplit[1:]
}
