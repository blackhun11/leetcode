package main

import "fmt"

func main() {
	fmt.Println(findTotalFactorEqual6FromN(128))
	fmt.Println(findTotalFactorEqual6FromN(1024))
	fmt.Println(findTotalFactorEqual6FromN(16384))
	fmt.Println(findTotalFactorEqual6FromN(262144))

}

func findTotalFactorEqual6FromN(n int) int {

	count := 0
	for i := 1; i <= n; i++ {
		if findFactorEqual6(i) {
			count++
		}
	}

	return count
}

func findFactorEqual6(in int) bool {

	factor := make(map[int]bool)
	lookup := make(map[int]bool)

	for i := 1; i <= in; i++ {

		if len(factor) > 6 {
			return false
		}

		if lookup[i] {
			continue
		}

		if in%i == 0 {
			factor[i] = true
			factor[in/i] = true
			lookup[in/i] = true
		}
	}

	if len(factor) != 6 {
		return false
	}

	return true
}
