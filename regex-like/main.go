package main

import (
	"flag"
	"github.com/teejays/clog"
)

const DEBUG bool = false

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

	run("ERROR: *.", "ERROR: file not found", true)
	run("ERROR: *.", "WARNING: file not found", false)

	// # an optional char that _can_ match is not forced to.
	// match "?aab" "ab"

	// # classic log searching
	// match "ERROR: *." "ERROR: file not found"
	// no_match "ERROR: *." "WARNING: file not found"
}

func run(pattern, str string, expected bool) {
	var answer bool = match(pattern, str)
	if answer == expected {
		clog.Infof("PASS  ==> Pattern: '%s' || String: '%s' || Answer: %t", pattern, str, answer)
	} else {
		clog.Warningf("FAIL  ==> Pattern: '%s' || String: '%s' || Answer: %t", pattern, str, answer)
	}
}

func match(_pattern, _str string) bool {

	var pattern []rune = []rune(_pattern)
	var str []rune = []rune(_str)

	// DP: build a memo table, where
	// T[i][j] represents whether pattern[:i]
	// matches string[:j]
	var memo [][]bool = make([][]bool, len(pattern)+1)

	if DEBUG {
		clog.Infof("Starting matching: pattern='%s', string='%s'", _pattern, _str)
	}

	for i := 0; i <= len(pattern); i++ {
		memo[i] = make([]bool, len(str)+1)

		for j := 0; j <= len(str); j++ {
			if DEBUG {
				//clog.Debugf("Memo: %v", memo)\
				clog.Debugf("Matching: pattern='%s', string='%s'", string(pattern[0:i]), string(str[0:j]))
				clog.Debugf("Populating Memo: i=%d, j=%d", i, j)
			}

			// Both pattern and string are empty
			if i == 0 && j == 0 {
				if DEBUG {
					clog.Debugf("-- empty pattern & str => %t", true)
				}
				memo[i][j] = true

			} else if i == 0 {
				// pattern is empty
				if DEBUG {
					clog.Debugf("-- empty pattern but non-empty string => %t", false)
				}
				memo[i][j] = false

			} else if j == 0 {
				// string is empty
				if DEBUG {
					clog.Debugf("-- empty string => ...")
				}
				// the pattern has even chars, and the previous pattern char was ? or *
				if i%2 == 0 && (pattern[i-2] == '?' || pattern[i-2] == '*') {
					if DEBUG {
						clog.Debugf("-- empty string but it's optional => ...")
					}
					memo[i][j] = memo[i-2][j]
				}

			} else if isSpecialRepeating(pattern[i-1]) {
				// if the current pattern char requires a following char
				if DEBUG {
					clog.Debugf("-- special character, go to next...")
				}

			} else if pattern[i-1] == str[j-1] || pattern[i-1] == '.' {
				// if the chars are same (works for first char as well)

				if DEBUG {
					clog.Debugf("-- current characters matched...")
				}
				// if previous was a special char
				if i > 1 && isSpecialRepeating(pattern[i-2]) {
					memo[i][j] = (memo[i][j-1] || memo[i-2][j-1])
				} else {
					memo[i][j] = memo[i-1][j-1]
				}

			} else {
				// if chars are not the same, and not special

				if i == 1 {

				} else if pattern[i-2] == '+' {
					// if previous pattern char is +

				} else if pattern[i-2] == '?' || pattern[i-2] == '*' {
					// if previous pattern char is ? or *
					memo[i][j] = memo[i-2][j]
				}
			}

			if DEBUG {
				if memo[i][j] {
					clog.Infof("Subpattern Match => %t", memo[i][j])
				} else {
					clog.Debugf("Subpattern Match => %t", memo[i][j])
				}
			}

		}
	}

	if DEBUG {
		clog.Debugf("Memo: %v", memo)
	}
	return memo[len(pattern)][len(str)]
}

func isSpecialRepeating(char rune) bool {
	if char == '?' || char == '*' || char == '+' {
		return true
	}
	return false
}
