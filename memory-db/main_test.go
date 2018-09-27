package main

import (
	"fmt"
	"github.com/teejays/clog"
	"testing"
)

type MemoryDbTest struct {
	Inputs  []string
	Outputs []string
}

var Example1 MemoryDbTest = MemoryDbTest{
	Inputs:  []string{"GET a", "SET a foo", "SET b foo", "COUNT foo", "COUNT bar", "DELETE a", "COUNT foo", "SET b baz", "COUNT foo", "GET b", "GET B", "END"},
	Outputs: []string{"NULL", "", "", "2", "0", "", "1", "", "0", "baz", "NULL", ""},
}

var Example2 MemoryDbTest = MemoryDbTest{
	Inputs:  []string{"SET a foo", "SET a foo", "COUNT foo", "GET a", "DELETE a", "GET a", "COUNT foo", "END"},
	Outputs: []string{"", "", "1", "foo", "", "NULL", "0", ""},
}

var Example3 MemoryDbTest = MemoryDbTest{
	Inputs:  []string{"BEGIN", "SET a foo", "GET a", "BEGIN", "SET a bar", "GET a", "ROLLBACK", "GET a", "ROLLBACK", "GET a", "END"},
	Outputs: []string{"", "", "foo", "", "", "bar", "", "foo", "", "NULL", ""},
}

var Example4 MemoryDbTest = MemoryDbTest{
	Inputs:  []string{"SET a foo", "SET b baz", "BEGIN", "GET a", "SET a bar", "COUNT bar", "BEGIN", "COUNT bar", "DELETE a", "GET a", "COUNT bar", "ROLLBACK", "GET a", "COUNT bar", "COMMIT", "GET a", "GET b", "END"},
	Outputs: []string{"", "", "", "foo", "", "1", "", "1", "", "NULL", "0", "", "bar", "1", "", "bar", "baz", ""},
}

func init() {
	EnableTestMode = true
	clog.LogLevel = 0
}

func TestInvalidCommand1(t *testing.T) {
	out, err := handleStatement("BLAH a foo")
	if err != ERR_INVALID_COMMAND_KEYWORD {
		t.Errorf("Expected ERR_INVALID_COMMAND_KEYWORD, but got: %s", err)
	}
}

func TestInvalidNumArgs(t *testing.T) {
	out, err := handleStatement("SET a foo bar")
	if err != ERR_INVALID_ARGS_NUM {
		t.Errorf("Expected ERR_INVALID_ARGS_NUM, but got: %s", err)
	}
}

func TestExample1(t *testing.T) {
	err := exampleTestHelper(Example1)
	if err != nil {
		t.Error(err)
	}
}

func TestExample2(t *testing.T) {
	err := exampleTestHelper(Example2)
	if err != nil {
		t.Error(err)
	}
}

func TestExample3(t *testing.T) {
	err := exampleTestHelper(Example3)
	if err != nil {
		t.Error(err)
	}
}

func TestExample4(t *testing.T) {
	err := exampleTestHelper(Example4)
	if err != nil {
		t.Error(err)
	}
}

func exampleTestHelper(test MemoryDbTest) error {
	if len(test.Inputs) != len(test.Outputs) {
		return fmt.Errorf("Invalid test params: length of inputs does not equal length of outputs")
	}
	clog.Debugf("ExampleTest Staring: Count %d \n", len(test.Inputs))

	for i := 0; i < len(test.Inputs); i++ {
		out, err := handleStatement(test.Inputs[i])
		clog.Infof("Statement #%d: %s | Expected: %s | Got: %s", i, test.Inputs[i], test.Outputs[i], out)
		if err != nil {
			return fmt.Errorf("Stmt: '%s' | Error: %s", test.Inputs[i], err)
		}
		if out != test.Outputs[i] {
			return fmt.Errorf("Stmt: '%s' | Expected out '%s', got '%s'", test.Inputs[i], test.Outputs[i], out)
		}

	}
	return nil
}
