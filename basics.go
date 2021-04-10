package main

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(2, 2) != 4 {
		t.Error("Expected 2+2 to equal 4")
	}
}

// Int function
func Addition(a int, b int) int {
	return a + b
}

// String forloop function
func StrForLoop(a string, b string) string {
	var strRes string = ""

	for i := 1; i < 5; i++ {
		strRes += a + " " + b + " "
	}

	return strRes

}

// While loop
func whilLoop() int {
	d := true
	n := 1
	for d {
		n *= 2
		if n > 5 {
			d = false
		}
	}
	return n
}

func userInput() {
	println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	print(first)
}

func main() {
	println("Hello World")
	println(Addition(25, 16))
	println(StrForLoop("Hello", "World"))
	println(whilLoop())
	userInput()
}
