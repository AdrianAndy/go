package main

import (
        "errors"
        "fmt"
        "os"
)

var (
    errorEmptyString = errors.New("Empty string. Try again.\n")
)
func printer(msg string) error {
    if msg == "" {
       panic(errorEmptyString)
    }
      _, err := fmt.Printf("%s\n", msg)

      return err
}

func main () {
    if err := printer(""); err != nil {
      if err == errorEmptyString {
          fmt.Printf("You tried to print an empty string!")
      } else {
        fmt.Printf("Printer failed: %s\n", err)
      }
        os.Exit(1)
    }
}
