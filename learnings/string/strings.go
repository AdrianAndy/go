  package main

import (
    "fmt"
)

func main () {
  atoz := "the quick brown fox jumps over the lazy dog\n"
// [0:9] - displays the first 9 characters. The first number selects how far in you want to start, second where to end.
// Leave blank to start at the start or go to the end.

  for i, r := range atoz {
     fmt.Printf("%d %c\n", i, r, len(atoz))
  }
}
