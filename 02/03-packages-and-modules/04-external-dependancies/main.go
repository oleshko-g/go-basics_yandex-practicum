package main

import (
	"fmt"

	ints "github.com/oleshko-g/math/Ints"
)

func main() {
	if sum := ints.AddInts(1, 2); sum != 3 {
		fmt.Printf("sum expected to be 3; got %d", sum)
	}

	fmt.Println("Well done!")
}
