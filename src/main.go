// Main entry point for a golang program
package main

import (
	"fmt"
	"runtime"
	"reflect"
)

// Alternate way to declare global variables
// var (
// 	name, course string
// 	module		 float64
// )

func main() {

	// All variables declared like this have to be used once, if inside a function
	// They can only be declared like this inside a funciton
	name := "Robert"
	// course := "Docker Deep Dive"
	module := 3.2


	fmt.Println("Hello, world.\n" + runtime.GOOS)
	fmt.Println("Name is set to", name, "and is of the type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and it is of type", reflect.TypeOf(module))

	// Pointers in GoLang
	fmt.Println("The memory adress of module is", &module)
}