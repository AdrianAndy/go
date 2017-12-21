package main

import (
	"fmt"

	"github.com/user/packages/shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	i := intSlice{3, 1, 4, 1, 5, 9}
	shuffler.Shuffle(i)
	fmt.Printf("%v", i)
}
