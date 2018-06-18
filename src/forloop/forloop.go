package main

import (
  "fmt"
)
//for loop with all conditions on one line
//func main() {
  //for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
    //fmt.Printf("%d あなたは悲しい兼を持っています\n", j)
  //}
//}


//bool

func main() {
  var stop bool

  i := 0

  for !stop {
    fmt.Printf("あなたは悲しい兼を持っています\n")

    i += 1
    if i == 10 {
      stop = true
    }
  }
}
