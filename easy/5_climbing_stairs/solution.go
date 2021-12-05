package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs/
func main() {
	fmt.Println(climbStairs(45))
	fmt.Println(climbStairs2(45))
}

// using recursive & memo (DP)
var memo [46]int

func climbStairs(N int) int {

	if N <= 2 {
		return N
	}

	if memo[N] != 0 {
		return memo[N]
	}

	totalStep := climbStairs(N-1) + climbStairs(N-2)
	memo[N] = totalStep
	return totalStep
}

// using iterative
func climbStairs2(N int) int {
	if N <= 2 {
		return N
	}

	first, second := 1, 2

	for i := 3; i < N; i++ {
		tempFirst := second
		second = first + second
		first = tempFirst
	}

	return first + second
}
