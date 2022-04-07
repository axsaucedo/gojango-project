package main

import (
	"fmt"

	"github.com/axsaucedo/gjango"
)

func main() {
	result := gjango.TestFunc(1, 1)
	fmt.Println(result)

	result = gjango.TestFunc2(1, 1)
	fmt.Println(result)

	result = gjango.TestFunc3(1, 2)
	fmt.Println(result)
}
