package main

import (
  "os/signal"
  "syscall"
  "fmt"
)

func main() {
  signal.Ignore(syscall.SIGTERM)
  signal.Ignore(syscall.SIGKILL)
  signal.Ignore(syscall.SIGINT)
  signal.Ignore(syscall.SIGABRT)
  signal.Ignore(syscall.SIGQUIT)

  for {
  	fmt.Printf("あなたは悲しい兼を持っています\n")
  	}

}
