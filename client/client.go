package main

import (
	"fmt"

	"github.com/thaJeztah/multimodule/internal/foo"
)


func main() {
	fmt.Println("hello client")
	fmt.Println(foo.Hello)
}
