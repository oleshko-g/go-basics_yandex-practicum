package main

import (
	"fmt"
	mathslice "mathSlice/internal/mathSlice"
)

func main() {
	s := mathslice.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("Sum slice: [%v].\n", mathslice.SumSlice(s))
	mathslice.MapSlice(s, func(i mathslice.Element) mathslice.Element { return i * 2 })
	fmt.Printf("Multiplied by 2 slice: [%v].\n", s)
	fmt.Printf("Sum slice: [%v].\n", mathslice.SumSlice(s))
	fmt.Printf("Folded by addition slice: [%v].\n",
		mathslice.FoldSlice(s,
			func(a mathslice.Element, b mathslice.Element) (c mathslice.Element) {
				return a + b
			}, 0))
	fmt.Printf("Folded by multiplication slice: [%v]",
		mathslice.FoldSlice(s,
			func(a mathslice.Element, b mathslice.Element) (c mathslice.Element) {
				return a * b
			}, 1))
}
