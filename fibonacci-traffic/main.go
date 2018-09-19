package main

import (
	"fmt"
	"time"
)

type TestParams struct {
	PastValues []int64
	N          int
	Expected   int64
}

func testAlgo(test TestParams) {
	result := "FAIL"
	start := time.Now()
	ans := algo(test.PastValues[0], test.PastValues[1], test.PastValues[2], test.N)
	if ans == test.Expected {
		result = "PASS"
	}
	fmt.Printf("*** %s *** Input: %d %d %d N: %d Expected: %d Got: %d | Time Taken: %s\n", result, test.PastValues[0], test.PastValues[1], test.PastValues[2], test.N, test.Expected, ans, time.Now().Sub(start).String())
}

var tests []TestParams = []TestParams{
	{[]int64{0, 0, 0}, 100, 0},

	{[]int64{50, 100, 150}, 1, 200},
	{[]int64{50, 100, 150}, 2, 300},
	{[]int64{50, 100, 150}, 3, 433},
	{[]int64{50, 100, 150}, 4, 622},
	{[]int64{50, 100, 150}, 5, 903},
	{[]int64{50, 100, 150}, 6, 1305},
	{[]int64{50, 100, 150}, 10, 5707},
	{[]int64{50, 100, 150}, 30, 9154511},
	{[]int64{50, 100, 150}, 90, 37787289279497674},

	{[]int64{1, 1, 2}, 30, 79453},
	{[]int64{1, 1, 2}, 100, 13135245745178652},
}

func main() {
	for _, test := range tests {
		testAlgo(test)
	}
}

func algo(day1, day2, day3 int64, n int) int64 {

	var cache []int64 = []int64{day1, day2, day3}

	for i := 3; i < n+3; i++ {
		cache = append(cache, 2*(cache[i-1]+cache[i-2]+cache[i-3])/3)
	}

	return cache[len(cache)-1]
}
