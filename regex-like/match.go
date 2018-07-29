package main

const (
	QUESTION_MARK rune = '?'
	ASTERISK           = '*'
	PLUS               = '+'
	DOT                = '.'
)

// Match implements a regex like pattern matcher using a Dynamic Programming
// approach. The following cases are implemented:
// 1) A non-special character in a pattern matches only that character.
// 2) The special-character `.` in the pattern matches any single character.
// 3) The special-character `?` in the pattern does not match any character, but
// indicates the following character in the pattern can match zero or one times.
// 4) The special-character `*` in the pattern does not match any character, but
//  indicates the following character in the pattern can match zero or more times.
// 5) The special-character `+` in the pattern does not match anything, but
//  indicates the following character in the pattern can match one or more times.
func Match(_pattern, _str string) bool {

	log(INFO, "Starting: pattern='%s', string='%s'", _pattern, _str)

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

		log(DEBUG, " - Starting loop")
		log(DEBUG, " - - Raw Indexes: i=%d, j=%d", i, j)

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

		log(DEBUG, " - - Cleaned Indexes: i=%d, j=%d", i, j)
		log(DEBUG, " - - Matching: pattern='%s', string='%s'", string(pattern[0:i]), string(str[0:j]))

		// LOGIC: there are multiple cases, each defining the recursive
		// relationship differently.

		if false { // only for formatting purposes

			// CASE 1: If both pattern and str are empty then it's a match.
			// We can save time by moving to the next pattern character
			// as an empty pattern would not match any str.
		} else if i == 0 && j == 0 {
			log(DEBUG, " - - - Empty pattern & str")
			memo[i][j] = true
			_i = i + 1

			// CASE 2: If the str is empty but pattern is not
		} else if j == 0 {
			log(DEBUG, " - - - Only str is empty")

			// CASE 2a: and the previous pattern char allows empty str e.g. '*c' allows empty
			// str, then we should use recursive relationship T[i][j] = T[i-2][j]
			// because if pattern[0:i] is '*a?b*c', then match os only possible if
			// pattern[0:i-2] is also matched '*a?b', and so on. We can then move to the next
			// iteration.
			if i%2 == 0 && (pattern[i-2] == QUESTION_MARK || pattern[i-2] == ASTERISK) {
				log(DEBUG, " - - - empty string but it's optional => ...")
				memo[i][j] = memo[i-2][j]
			}
			_j = j + 1 // next str character

			// CASE 3: If the last character of the current sub-pattern 'requires'
			// a following character, then the current sub-pattern is invalid and
			// would never match. e.g. 'abc*' is invalid as it requires another character.
			// We can save time by moving to the next pattern character, as such pattern
			// would never match any str.
		} else if isAmong(pattern[i-1], QUESTION_MARK, ASTERISK, PLUS) {
			log(DEBUG, " - - - pattern character is special & requires a following character...")
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
			log(DEBUG, " - - - last characters matched...")

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
			log(INFO, " - - Subpattern Match => %t", memo[i][j])
		} else {
			log(NOTICE, " - - Subpattern Match => %t", memo[i][j])
		}

		// set the pointers for next iteration to the new values
		i = _i
		j = _j

	}

	return memo[len(pattern)][len(str)]
}

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
