package main

import (
        "fmt"
        "time"
)

func emit(wordChannel chan string, done chan bool){
      words := []string{"The", "Quick", "Brown", "Fox"}
      defer close(wordChannel)

      i := 0
      t := time.NewTimer(3 * time.Second)

      for {
          select {
          case wordChannel <- words[i]:
            i += 1
            if i == len(words) {
              i = 0
            }
          case <-done:
            done <- true
              return

            case <-t.C:
              return
          }
      }
}

func main(){
    wordCh := make(chan string)
    doneCh := make(chan bool)

    go emit(wordCh, doneCh)

    for i := 0; i < 100; i++ {
          fmt.Printf("%s ", <-wordCh)
    }

    doneCh <- true
    <-doneCh
}
