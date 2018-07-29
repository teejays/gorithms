package main

import (
	"fmt"
	"github.com/teejays/clog"
	"testing"
)

type MatchTestCase struct {
	Pattern  string
	Str      string
	Expected bool
}

var testCases []MatchTestCase

func init() {
	testCases = []MatchTestCase{
		{Pattern: "abc", Str: "abc", Expected: true},
		{Pattern: "abd", Str: "abc", Expected: false},
		{Pattern: "a.c", Str: "a.c", Expected: true},
		{Pattern: "a.c", Str: "abc", Expected: true},
		{Pattern: "a?bc", Str: "abc", Expected: true},
		{Pattern: "a?bc", Str: "ac", Expected: true},
		{Pattern: "?aab", Str: "ab", Expected: true},
		{Pattern: "a+b*b", Str: "ab", Expected: true},
		{Pattern: "a+b*b", Str: "abbbb", Expected: true},
		{Pattern: "a+b+b", Str: "ab", Expected: false},
		{Pattern: "a*b+b", Str: "abbbb", Expected: true},
		{Pattern: "a*b*b", Str: "a", Expected: true},
		{Pattern: "ERROR: *.", Str: "ERROR: file not found", Expected: true},
		{Pattern: "ERROR: *.", Str: "WARNING: file not found", Expected: false},
	}
}

func TestMatch(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), testRun(c.Pattern, c.Str, c.Expected))
	}
}

func testRun(pattern, str string, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		answer := Match(pattern, str)
		clog.Infof("Pattern: '%s' || String: '%s' => Answer: %t", pattern, str, answer)
		if answer != expected {
			clog.Warning("Wrong answer!")
			t.Error("Wrong answer!")
		}
	}
}
