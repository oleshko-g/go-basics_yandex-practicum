package main

import (
	"fmt"

	"go-basics_yandex-practicum/02/03-packages-and-modules/03-modules/somemodule/somepackage"
	"ypmodule/calc"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))
	somepackage.Func()
}
