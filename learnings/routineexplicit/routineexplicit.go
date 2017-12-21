package main

import (
        "fmt"
)

//concurrent go routine emitting via word channel, word by word
//this version is explicit
func makeID(idChan chan int) {
  var id int
  id = 0

  for {
      idChan <- id
      id += 1
  }
}

//receiving from the wordchannel, word by word.
//func main() {
  //  wordChannel := make(chan string)

    //go emit(wordChannel)

    //word := <-wordChannel
    //fmt.Printf("%s\n", word)
    //word = <-wordChannel
    //fmt.Printf("%s\n", word)
    //word = <-wordChannel
    //fmt.Printf("%s\n", word)
    //word = <-wordChannel
    //fmt.Printf("%s\n", word)
    //word, ok := <-wordChannel
    //fmt.Printf("%s %t\n", word, ok)
//}

//receiving from the wordchannel, word by word.
func main() {
    idChan := make(chan int)

    go makeID(idChan)
    fmt.Printf("%d ", <-idChan)
    fmt.Printf("%d ", <-idChan)
    fmt.Printf("%d ", <-idChan)
    }
