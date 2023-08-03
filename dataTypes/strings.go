package datatypes

import "fmt"

func Strings() {
	// Declare a string
	str := "Hello, Gophers!"

	// Attempt to modify the string directly (which is NOT allowed)
	// Uncommenting the line below will result in a compile-time error.
	// str[0] = 'h'

	// Instead, to modify the string, we create a new one with the desired changes.
	newStr := "hello" + str[5:]

	// Original string remains unchanged
	fmt.Println("Original string:", str)

	// Modified string
	fmt.Println("Modified string:", newStr)
}

// In the example, we have a string str initialized with the value "Hello, Gophers!".
// If we try to modify the string directly by assigning a new character to an index, like str[0] = 'h', it will result in a compile-time error because strings are immutable.
