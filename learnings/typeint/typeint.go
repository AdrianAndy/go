package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}
func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := stringSlice{"I", "wish", "the", "milkman", "would", "deliver", "my", "milk", "in", "the", "morning"}
	shuffle(s)
	fmt.Printf("%v\n", s)
}
