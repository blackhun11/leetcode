package main

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter2([]byte{'c', 'f', 'j'}, 'a')))

}

// my original code, still o(nlogn)
func nextGreatestLetter(letters []byte, target byte) byte {

	length := len(letters) - 1
	left := 0
	right := length
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if letters[mid] == target {
			if length == mid {
				return letters[0]
			}
			if letters[mid+1] == target {
				left = mid + 2
				continue
			}
			return letters[mid+1]

		} else if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if length == mid && letters[mid] < target {
		return letters[0]
	}

	if left > length {
		if letters[length] <= target {
			return letters[0]
		} else {
			return letters[length]
		}
	}

	return letters[left]
}

// cheating answer
func nextGreatestLetter2(letters []byte, target byte) byte {

	left := 0
	right := len(letters)

	for left < right {
		mid := (left + right) / 2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
		fmt.Println(left, right)
	}

	if left >= len(letters) {
		return letters[0]
	}

	return letters[left]
}
