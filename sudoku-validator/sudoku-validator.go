// Sudoku Validator
//
// Sudoku puzzles are 9x9 logic puzzles akin to crosswords where the goal of the solver is to fill in the empty spaces with numbers 1-9 such that the puzzle is "valid".
//
// Have a look at an example Sudoku puzzle:
//      - https://upload.wikimedia.org/wikipedia/commons/e/e0/Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg
//
//This puzzle has a possible solution:
//      - https://upload.wikimedia.org/wikipedia/commons/1/12/Sudoku_Puzzle_by_L2G-20050714_solution_standardized_layout.svg
//
// There are 3 simple rules to solving a Sudoku puzzle.
//       1.  Each row in the puzzle is exactly 1 through 9. No duplicates.
//       2.  Each column in the puzzle is exactly 1 through 9.  No duplicates.
//       3.  Each 3x3 matrix in the puzzle is exactly 1 through 9. No duplicates.
//
// Part 1.
//   Write a method that will validate the rows of the puzzle
//
// Part 2.
//   Write a method that will validate the columns of the puzzle
//
// Part 3.
//   Write a method that will validate the nine 3x3 matrixes of the puzzle
//
// To execute Go code, please declare a func main() in a package "main"
package main

import "fmt"

var (
	goodPuzzle = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	badPuzzle = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 4},
		{6, 7, 2, 1, 90, 5, 3, 4, -8},
		{3, 9, -8, 3, 4, 2, 5, 6, 7625},
		{8, 5, 9, 6, 6, 1, 4, 2, 3},
		{4, 20, 6, 0, 5, 3, 7, -9, 1},
		{7, 1, 31, 9, 2, 4, 8, 5, 6},
		{9, 6, 21, 5, 3, 7, 2, 8, 4},
		{2, 6, 7, 2, 6, 8, 4, -13, 5},
		{3, 4, 5, 2, 8, 0, 1, 7, 9},
	}
)

const LOG_ERROR = false

var numTests, numPass int

func main() {

	// Test Row Validation
	fmt.Println("Testing Row Validation")

	for i := 0; i < 9; i++ {
		assertValidation(validateRow, goodPuzzle, i, true)
	}

	for i := 0; i < 9; i++ {
		assertValidation(validateRow, badPuzzle, i, false)
	}

	// Test Column Valdation
	fmt.Println("Testing Column Validation")

	for i := 0; i < 9; i++ {
		assertValidation(validateColumn, goodPuzzle, i, true)
	}

	for i := 0; i < 9; i++ {
		assertValidation(validateColumn, badPuzzle, i, false)
	}

	// Test Matrix Validation
	for i := 0; i < 9; i++ {
		assertValidation(validateMatrix, goodPuzzle, i, true)
	}
	for i := 0; i < 9; i++ {
		assertValidation(validateMatrix, badPuzzle, i, false)
	}

	// Summary
	fmt.Printf("\n Summary: Total Tests : %d | Tests Passed: %d \n", numTests, numPass)
	if numTests == numPass {
		fmt.Println("******** PASS ********")
	} else {
		fmt.Println("******** FAIL ********")
	}

}

func assertValidation(f func([][]int, int) (bool, error), puzzle [][]int, index int, expected bool) {
	numTests++

	result, err := f(puzzle, index)
	var message string = "FAIL"
	if result == expected {
		numPass++
		message = "PASS"
	}
	fmt.Printf("**** %s **** Expected: %v | Got: %v\n", message, expected, result)
	if err != nil && LOG_ERROR {
		fmt.Printf("**** Err **** %s\n", err)
	}

}

func validateRow(puzzle [][]int, rowNum int) (bool, error) {
	return validateSet(puzzle[rowNum])
}

func validateColumn(puzzle [][]int, colIndex int) (bool, error) {
	if len(puzzle) > 9 {
		return false, fmt.Errorf("Invalid length of puzzle: %d", len(puzzle))
	}

	var set []int

	for _, row := range puzzle {
		for j, val := range row {
			if j == colIndex {
				set = append(set, val)
			}
		}
	}

	return validateSet(set)

}

func validateMatrix(puzzle [][]int, matrixNum int) (bool, error) {

	const numRowsMatrix int = 3
	const numColsMatrix int = 3
	// MatrixNumber will be from 0-8, corresponding to the matric number in the puzzle
	// we need to find out what number/

	var matrixRowNum int = matrixNum / numRowsMatrix // so, matrix 0,1,2 / 3 => 0 || 3,4,5 => 1 || 7,8,9 => 3
	var matrixColNum int = matrixNum % numColsMatrix // so matrix 0,3,6 => 0 || 1,4,7 => 1 || 2,5,8 => 3

	var set []int
	var rowStart int = matrixRowNum * 3
	var colStart int = matrixColNum * 3

	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			set = append(set, puzzle[i][j])
		}
	}

	return validateSet(set)

}

func validateSet(set []int) (bool, error) {
	var seen []int = make([]int, 10)
	if len(set) != 9 {
		return false, fmt.Errorf("Unexpected length of set recieved: %d", len(set))
	}
	for _, i := range set {
		if i > 9 {
			return false, fmt.Errorf("a number in the set is greater than 9: %d", i)
		}
		if i < 1 {
			return false, fmt.Errorf("a number in the set is last than 1: %d", i)
		}
		if seen[i] == i {
			return false, fmt.Errorf("a number appears more than twice in the set: %d", i)
		}
		seen[i] = i
	}

	return true, nil

}

