package main

import (
	"fmt"

	"github.com/user/poetry"
)

func main() {
	p, err := poetry.LoadPoem("yermum")
	if err != nil {
		fmt.Printf("Error loading file: %s", err)
	}
	fmt.Printf("%s\n", p)
}
