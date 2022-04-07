package main

import (
	"fmt"

	"github.com/axsaucedo/gjango"
)

func main() {
	result := gjango.TestFunc(1, 1)
	gjango.TestFunc()

	fmt.Println(result)
}
