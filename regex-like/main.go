package main

import (
	"flag"
	"github.com/teejays/clog"
	"time"
)

func match(_pattern, _str string) bool {

	var pattern []rune = []rune(_pattern)
	var str []rune = []rune(_str)

	// DP: build a memo table, where
	// T[i][j] represents whether pattern[:i]
	// matches string[:j]
	var memo [][]bool = make([][]bool, len(pattern)+1)

	print(INFO, "Starting matching: pattern='%s', string='%s'", _pattern, _str)

	var i, j, _i, _j int
	for {

		print(DEBUG, "Raw Indexes: i=%d, j=%d", i, j)

		if i > len(pattern) {
			break
		}
		if j > len(str) {
			i++
			j = 0
			continue
		}
		_i, _j = i, j

		if memo[i] == nil {
			memo[i] = make([]bool, len(str)+1)
		}

		print(DEBUG, "Cleaned Indexes: i=%d, j=%d", i, j)
		print(DEBUG, "Matching: pattern='%s', string='%s'", string(pattern[0:i]), string(str[0:j]))

		// Both pattern and string are empty
		if i == 0 && j == 0 {
			print(DEBUG, "-- empty pattern & str")
			memo[i][j] = true
			_i = i + 1

		} else if i == 0 {
			// pattern is empty
			print(DEBUG, "SHOULD NOT HAPPEN -- empty pattern but non-empty string => %t", false)
			memo[i][j] = false
			_j = j + 1

		} else if j == 0 {
			// string is empty
			print(DEBUG, "-- empty string")

			// the pattern has even chars, and the previous pattern char was ? or *
			if i%2 == 0 && (pattern[i-2] == '?' || pattern[i-2] == '*') {
				print(DEBUG, "-- empty string but it's optional => ...")
				memo[i][j] = memo[i-2][j]
			}
			_j = j + 1

		} else if pattern[i-1] == '?' || pattern[i-1] == '*' || pattern[i-1] == '+' {
			// if the current pattern char requires a following char
			print(DEBUG, "-- special character, go to next...")
			_i = i + 1
			_j = 0

		} else if pattern[i-1] == str[j-1] || pattern[i-1] == '.' {
			// if the chars are same (works for first char as well)
			print(DEBUG, "-- current characters matched...")

			// if previous was a special char
			if i > 1 && (pattern[i-2] == '?' || pattern[i-2] == '*') {
				memo[i][j] = (memo[i][j-1] || memo[i-2][j-1] || memo[i-2][j])
			} else if i > 1 && pattern[i-2] == '+' {
				memo[i][j] = (memo[i][j-1] || memo[i-2][j-1])
			} else {
				memo[i][j] = memo[i-1][j-1]
			}
			_j = j + 1

		} else {
			// if chars are not the same, and not special
			if i > 1 && pattern[i-2] == '+' {
				// if previous pattern char is +

			} else if i > 1 && (pattern[i-2] == '?' || pattern[i-2] == '*') {
				// if previous pattern char is ? or *
				memo[i][j] = memo[i-2][j]
			}
			_j = j + 1
		}

		if memo[i][j] {
			print(INFO, "Subpattern Match => %t", memo[i][j])
		} else {
			print(DEBUG, "Subpattern Match => %t", memo[i][j])
		}

		i = _i
		j = _j

	}

	print(DEBUG, "Memo: %v", memo)
	return memo[len(pattern)][len(str)]
}

func isSpecialRepeating(char rune) bool {
	if char == '?' || char == '*' || char == '+' {
		return true
	}
	return false
}

const PRINT bool = false

const (
	DEBUG uint = iota
	INFO
)

func print(logger uint, message string, args ...interface{}) {
	if PRINT {
		switch logger {
		case DEBUG:
			clog.Debugf(message, args...)
		case INFO:
			clog.Infof(message, args...)
		}

	}
}

func main() {
	clog.GetCloggerByName("Debug").AddDecoration(clog.FG_YELLOW)
	var pattern, str string
	flag.StringVar(&pattern, "pattern", "", "the pattern used to match the string")
	flag.StringVar(&str, "string", "", "the string which is matched using the pattern")
	flag.Parse()

	if pattern != "" || str != "" {
		match(pattern, str)
	}

	run("abc", "abc", true)
	run("abd", "abc", false)

	run("a.c", "a.c", true)
	run("a.c", "abc", true)

	run("a?bc", "abc", true)
	run("a?bc", "ac", true)

	run("?aab", "ab", true)

	run("a+b*b", "ab", true)
	run("a+b*b", "abbbb", true)
	run("a+b+b", "ab", false)
	run("a*b+b", "abbbb", true)
	run("a*b*b", "a", true)

	run("ERROR: *.", "ERROR: file not found", true)
	run("ERROR: *.", "WARNING: file not found", false)
}

func run(pattern, str string, expected bool) {
	start := time.Now()
	var answer bool = match(pattern, str)
	elapsed := time.Now().Sub(start)

	if answer == expected {
		clog.Infof("***PASS*** ==> Pattern: '%s' || String: '%s' || Answer: %t || Time: %s", pattern, str, answer, elapsed)
	} else {
		clog.Warningf("***FAIL***  ==> Pattern: '%s' || String: '%s' || Answer: %t || Time: %s", pattern, str, answer, elapsed)
	}
}
