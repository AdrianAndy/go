package main

import (
  "fmt"
)

func printer(w []string) {
  for _, word := range w {
    fmt.Printf("%s", word)
    }
    fmt.Printf("\n")
}

func main() {
    words := make([]string, 0, 4)
    words = append(words, "The")
    words = append(words, "fucking")
    words = append(words, "brown")
    words = append(words, "fox")
    printer(words)

    newWords := make([]string, 4)
    copy(newWords, words) //use copy when you actually need to copy data, but it's mostly unused as we want to work with one set
    newWords[2] = "blue"
    printer(newWords)
}
