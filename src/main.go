// Main entry point for a golang program
package main

import (
	"fmt"
	// "os"
	// "runtime"
	// "reflect"
)

// Alternate way to declare global variables
// var (
// 	name, course string
// 	module		 float64
// )

func main() {

	var average = avergeScores(75, 33, 82, 99, 0, 100, 55);

	fmt.Println(average)

	// All variables declared like this have to be used once, if inside a function
	// They can only be declared like this inside a funciton
	// name := os.Get
	// course := "Docker Deep Dive"
	// // module := 3.2

	// fmt.Println("\nHi", name, "you're currently watching", course)

	// changeCourse(course)

	// fmt.Println("\nYou are now watching course", course)

// 	fmt.Println("Hello, world.\n" + runtime.GOOS)
// 	fmt.Println("Name is set to", name, "and is of the type", reflect.TypeOf(name))
// 	fmt.Println("Module is set to", module, "and it is of type", reflect.TypeOf(module))

// 	// Pointers in GoLang
// 	fmt.Println("The memory adress of module is", &module)
}

// Returns a string
func changeCourse(course string) string {
	course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)

	return course

}

// Variadic functions that do not need to know how many passed parameters there are.
func avergeScores(scores ...int) int {
	totalScore := 0
	count := 0
	for _, i := range scores {
		totalScore = totalScore + i
		count++	
	}
	return totalScore / count
}