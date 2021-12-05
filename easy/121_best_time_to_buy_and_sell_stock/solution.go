package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	str := readGist()
	in := toArrayofInt(str)

	fmt.Println(maxProfit(in))
	fmt.Println(maxProfit([]int{5, 4, 3, 2, 1}))
	fmt.Println(maxProfit([]int{3, 2, 1, 5, 6, 2}))

}

func readGist() string {
	response, err := http.Get("https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return buf.String()
}

func toArrayofInt(in string) []int {
	split := strings.Split(in, " ")

	out := make([]int, len(split))

	for _, v := range split {
		vInt, _ := strconv.Atoi(v)
		out = append(out, vInt)
	}

	return out
}

func maxProfit(prices []int) int {

	min := math.MaxInt16
	profit := 0

	for i := 0; i < len(prices); i++ {
		curPrice := prices[i]
		if curPrice < min {
			min = curPrice
			continue
		}

		curProfit := curPrice - min

		if curProfit > profit {
			profit = curProfit
		}
	}

	return profit
}

// brute force
func maxProfit2(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		buy := prices[i]

		for j := i + 1; j <= len(prices)-1; j++ {
			sell := prices[j]

			if sell-buy > profit {
				profit = sell - buy
			}
		}
	}

	return profit
}
