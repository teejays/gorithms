// main package implements a regex-like pattern matching
// algorithm using the instructions below in this
// coderpad. The algorithm takes a dynamic programming
// approach to pattern matching.
//
// The code is intended to err towards the side of
// over-commenting in order to communicate the thought
// process as well as possible.
//
// A version of the implementation is also available at:
// github.com/teejays/gorithms/tree/master/regex-like
package main

import (
	"fmt"
	"time"
)

// LOG const determines whether the program logs to the Std. Output
const LOG bool = false

/**************************************************
* I N S T R U C T I O N S
***************************************************/

/*
Your previous Plain Text content is preserved below:

# Implementing a Regular-Expression Searcher

For this exercise, create a simple pattern matcher (similar-to but
different-than regular expressions) that takes two arguments:

- First, the pattern
- Second, the string you want to match

If the pattern matches the string as defined below, return true.
Otherwise, return false.

## Matching Behavior

Your application should support these patterns:

- A non-special character in a pattern matches only that character.
- The special-character `.` in the pattern matches any single character.
- The special-character `?` in the pattern does not match any character, but
  indicates the following character in the pattern can match zero or one times.
- The special-character `*` in the pattern does not match any character, but
  indicates the following character in the pattern can match zero or more times.
- The special-character `+` in the pattern does not match anything, but
  indicates the following character in the pattern can match one or more times.

### Details

The pattern must match every character in the string to be considered a match -
we are only matching strings in their entirety, unlike `grep` or similar.

For example, the string `abc` matches the patterns `abc`, `...`, `a.c`, and
`a?bbc` but does _not_ match `..` or `....` (since it must match completely).

You should read the above definitions and then add appropriate test cases in
addition to those described below.

Note: Do not worry about escaping special characters (e.g., `+`).

## Testing

Here are a few test cases to get you started (though you will certainly want to
add more):
```
# exact match and simple mismatch
match "abc" "abc"
no_match "abd" "abc"

# any-char matches
match "a.c" "a.c"
match "a.c" "abc"

# an optional pattern char matches with and without
match "a?bc" "abc"
match "a?bc" "ac"

# an optional char that _can_ match is not forced to.
match "?aab" "ab"

# classic log searching
match "ERROR: *." "ERROR: file not found"
no_match "ERROR: *." "WARNING: file not found"
```


## Submitting Your Solution

Use any language you're comfortable with and feel free to consult any
documentation, StackOverflow, etc. as you normally would in your day-to-day
work.

Pick your desired language from the drop-down above, or, if it isn't supported by
CoderPad, just let us know and include instructions on how to build and run your
solution on recent-ish Linux or MacOS systems.

You can develop and run your solution right in CoderPad if you want, or work in
your preferred environment and paste your code in when you're ready -- just make
sure that it does run in CoderPad before you're done. Once you're finished, just
let us know!


Note: Since the goal is to _implement_ a form of matching, please _do not_ use
regex libraries in your solution.
*/

/**************************************************
* A L G O R I T H M
***************************************************/

const (
	QUESTION_MARK rune = '?'
	ASTERISK           = '*'
	PLUS               = '+'
	DOT                = '.'
)

// Match implements a regex like pattern matcher using
// dynamic Programming approach.
func Match(_pattern, _str string) bool {

	log("Starting: pattern='%s', string='%s'", _pattern, _str)

	// Rune: Go treats characters as type rune. A rune is like a
	// []byte, where one or more bytes combine to form a character.
	// For efficient handling, convert the strings to []rune
	var pattern []rune = []rune(_pattern)
	var str []rune = []rune(_str)

	// Dynamic Programming:
	// T[i][j] represents whether sub-pattern[0:i] matches sub-str[0:j].
	// Depending on the last character of the sub-pattern & sub-str,
	// T[i][j] can depend on T[i-1][j-1], or T[i-2][j], or T[i-2][j-1]

	// We can create a table that can store T[i][j] values as go through
	// the pattern and the str
	var memo [][]bool = make([][]bool, len(pattern)+1)
	for n, _ := range memo {
		memo[n] = make([]bool, len(str)+1)
	}

	// i, j are the main pointers that point to i & j in T[i][j] (described above).
	var i, j int

	// _i, _j are tempt variables that hold the next (upcoming) i, j values.
	// We store them in temp variables because we don't want to influence original i & j
	// as we go through the loop
	var _i, _j int

	// Start the loop, with i & j initialized to zero, and dynamically change the
	// i & j value as we progress until we reach the end.
	for {

		log(" - Starting loop")
		log(" - - Raw Indexes: i=%d, j=%d", i, j)

		// Base condition for breaking the loop
		if i > len(pattern) {
			break
		}
		// this mimics kind of a nested for-loop. If j reaches the limit,
		// increment i and set j to 0, as that's the next iteration
		if j > len(str) {
			i++
			j = 0
			continue
		}

		// update the temp pointers to reflect the actual pointer values
		// _i, _j are supposed to mirror i & j at the start of the iteration
		_i, _j = i, j

		log(" - - Cleaned Indexes: i=%d, j=%d", i, j)
		log(" - - Matching: pattern='%s', string='%s'", string(pattern[0:i]), string(str[0:j]))

		// LOGIC: there are multiple cases, each defining the recursive
		// relationship differently.

		if false { // only for formatting purposes

			// CASE 1: If both pattern and str are empty then it's a match.
			// We can save time by moving to the next pattern character
			// as an empty pattern would not match any str.
		} else if i == 0 && j == 0 {
			log(" - - - Empty pattern & str")
			memo[i][j] = true
			_i = i + 1

			// CASE 2: If the str is empty but pattern is not
		} else if j == 0 {
			log(" - - - Only str is empty")

			// CASE 2a: and the previous pattern char allows empty str e.g. '*c' allows empty
			// str, then we should use recursive relationship T[i][j] = T[i-2][j]
			// because if pattern[0:i] is '*a?b*c', then match os only possible if
			// pattern[0:i-2] is also matched '*a?b', and so on. We can then move to the next
			// iteration.
			if i%2 == 0 && (pattern[i-2] == QUESTION_MARK || pattern[i-2] == ASTERISK) {
				log(" - - - empty string but it's optional => ...")
				memo[i][j] = memo[i-2][j]
			}
			_j = j + 1 // next str character

			// CASE 3: If the last character of the current sub-pattern 'requires'
			// a following character, then the current sub-pattern is invalid and
			// would never match. e.g. 'abc*' is invalid as it requires another character.
			// We can save time by moving to the next pattern character, as such pattern
			// would never match any str.
		} else if isAmong(pattern[i-1], QUESTION_MARK, ASTERISK, PLUS) {
			log(" - - - pattern character is special & requires a following character...")
			_i = i + 1 // next pattern character
			_j = 0

			// CASE 4: If the last character of pattern matches the last character
			// of str (same character or DOT)
			// Three different recursive relationships can apply here:
			// 1) The current sub-str character is the first character to satisfy the
			// current pattern character. This means that pattern minus the
			// current pattern character i.e. pattern[0:i-2] satisfies the sub-str
			// without the current char i.e. str[0:j-1].
			// e.g. 'a(?|*|+)b' & 'ab'
			// 2) The current str character is not the first character satisfying the current
			// pattern character. same sub-pattern i.e. pattern[O:i], also matched the
			// sub-str without the current char i.e. str[0:j-1].
			// e.g 'a(?|*|+)b' & 'abbb'
			// 3) This current sub-str satisified the previous pattern character as well.
			// This means that pattern minus the current pattern character i.e.
			// pattern[0:i-2] is satisfied by the current sub-str i.e. sur[0:j].
			// e.g. 'a*b(?|*)b' & 'abb'
		} else if isAmong(pattern[i-1], str[j-1], DOT) {
			log(" - - - last characters matched...")

			// CASE 4a: and the previous pattern character was such that it is not
			// neccesary to have a string character to match it e.g. (QUESTION_MARK, ASTERISK)
			// then the sub-pattern and sub-str will match in three cases:

			if i > 1 && isAmong(pattern[i-2], QUESTION_MARK, ASTERISK) {
				memo[i][j] = (memo[i-2][j-1] || memo[i][j-1] || memo[i-2][j])

				// CASE 4b: and previous patter character was PLUS,
				// then it will match in cases 1 & 2:
				// NOTE: 3) The third case from 4a will not apply here because it is neccesary'
				// for the current sub-str character to match the current pattern character
			} else if i > 1 && isAmong(pattern[i-2], PLUS) {
				memo[i][j] = (memo[i-2][j-1] || memo[i][j-1])

				// CASE 4c: and previous character has no relevance to current. Only case
				// 3 applies here.
			} else {
				memo[i][j] = memo[i-1][j-1]
			}

			_j = j + 1 // next str character

			// CASE 5: If the last characters do not match, and are not special
		} else {

			// CASE 5b: if previous pattern character allowed the current pattern
			// character to not exist, then it's okay, we move on to next pattern
			if i > 1 && isAmong(pattern[i-2], QUESTION_MARK, ASTERISK) {
				memo[i][j] = memo[i-2][j]
			}

			_j = j + 1 // next str character
		}

		if memo[i][j] {
			log(" - - Subpattern Match => %t", memo[i][j])
		} else {
			log(" - - Subpattern Match => %t", memo[i][j])
		}

		// set the pointers for next iteration to the new values
		i = _i
		j = _j

	}

	return memo[len(pattern)][len(str)]
}

/**************************************************
* U T I L  F U N C T I O N S
***************************************************/

// isAmong takes a rune and bunch of rune args and returns true if the
// rune matches any of the args
func isAmong(r rune, args ...rune) bool {
	for _, arg := range args {
		if r == arg {
			return true
		}
	}
	return false
}

func log(message string, args ...interface{}) {
	if LOG {
		fmt.Printf(message+"\n", args...)
	}
}

/**************************************************
* M A I N
***************************************************/

func main() {
	// Get all the tests & run them
	testCases := getTestCases()
	for _, t := range testCases {
		runTest(t)
	}
}

/**************************************************
* T E S T S
***************************************************/
// TestCase type represents all params needed to
// succesfully run a test for this algorithm
type TestCase struct {
	Pattern  string
	Str      string
	Expected bool
}

// getTestCases defines and returns the tests that will
// run upon execution. More tests can be added here.
func getTestCases() []TestCase {
	var testCases []TestCase

	testCases = []TestCase{
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

	return testCases
}

// RunTest executes a given test case, logging info on
// whether it passed or failed.
func runTest(t TestCase) {
	answer := run(t.Pattern, t.Str)
	result := "FAIL"
	if answer == t.Expected {
		result = "PASS"
	}
	fmt.Printf("***%s***\n\n", result)
}

// run calls the algorithm (Match) and logs the results.
func run(pattern, str string) bool {
	start := time.Now()
	answer := Match(pattern, str)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Pattern: '%s' || String: '%s' ||  Time: %s => Answer: %t \n", pattern, str, elapsed, answer)
	return answer
}
