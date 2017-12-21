package main

import (
        "fmt"
)

//concurrent go routine emitting via word channel, word by word
//this version is using range
func emit(c chan string) {
    words := []string{"The", "Quick", "Brown", "Fox"}

    for _, word := range words {
      c <- word
    }

    close(c)
}

//receiving from the wordchannel, word by word.
func main() {
    wordChannel := make(chan string)

    go emit(wordChannel)

    for word := range wordChannel {
        fmt.Printf("%s ", word)
    }

    fmt.Printf("\n")
}
