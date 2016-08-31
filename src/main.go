package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello, world.\n" + runtime.GOOS)
}