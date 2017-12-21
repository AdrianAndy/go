package main

import (
	"fmt"
)

func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case uint32:
		fmt.Printf("It's an unassigned 32-bit integer: %d\n", v)
	default:
		fmt.Printf("Dunno, mate: %v\n", v)
	}
}

func main() {
	whatIsThis([...]string{"A", "B", "C"})
	//whatIsThis(uint32(42))
	//whatIsThis("Hello")
}
