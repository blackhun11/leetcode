package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// str := readGist()
	// fmt.Println(uniqueFirstOccurence(str))

	// fmt.Println(uniqueFirstOccurence("sebaerb"))
	fmt.Println(lexiographical("cbacdcbc"))
}

func readGist() string {
	response, err := http.Get("https://gist.githubusercontent.com/Jekiwijaya/0b85de3b9ff551a879896dd78256e9b8/raw/e9d58da5d4df913ad62e6e8dd83c936090ee6ef4/gistfile1.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return buf.String()
}

func uniqueFirstOccurence(str string) string {

	lookup := make(map[byte]bool)

	byteIn := []byte(str)

	out := []byte{}

	for _, v := range byteIn {

		if lookup[v] {
			continue
		}

		out = append(out, v)
		lookup[v] = true
	}
	return string(out)
}

// https://www.youtube.com/watch?v=0tantogp8fc
func lexiographical(str string) string {
	byteVal := []byte(str)

	count := make(map[byte]int)

	stack := stack{}
	added := make(map[byte]bool)

	for _, v := range byteVal {
		count[v]++
	}

	for _, v := range byteVal {
		count[v]--

		if added[v] {
			continue
		}

		for !stack.isEmpty() && stack.getLast() > v && count[stack.getLast()] > 0 {
			delete(added, stack.pop())
		}

		stack.push(v)
		added[v] = true
	}
	return string(stack)
}

type stack []byte

func (s *stack) push(b byte) {
	*s = append(*s, b)
}

func (s *stack) getLast() byte {
	arr := *s
	return arr[len(arr)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop() byte {
	var val byte
	if len(*s) > 0 {
		arr := *s
		val = arr[len(arr)-1]
		*s = arr[:len(arr)-1]
	}
	return val
}
