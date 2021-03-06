// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Create a package named toy with a single unexported struct type named bat. Add
// the exported fields Height and Weight to the bat type. Then create an exported
// factory method called NewBat that returns pointers of type bat that are initialized
// to their zero value.
//
// Create a program that imports the toy package. Use the NewBat function to create a
// value of bat and populate the values of Height and Width. Then display the value of
// the bat variable.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Create a value of type bat.
	variable_name := toy.function_name()
	variable_name.field_name = value
	variable_name.field_name = value

	// Display the value.
	fmt.Println(variable_name)
}
