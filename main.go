/*
// GO Constants
package main

import (
	"fmt"
)

const FNAME, LNAME = "Josiah", "Olaniyi"
const AGE int = 21
const GENDER string = "Male"
const COMP string = "Dark"

func main() {
	fmt.Println(FNAME, "", LNAME) // making the name on a single line
	fmt.Println(AGE)
	fmt.Println(GENDER)
	fmt.Println(COMP)
} 

// Making a new line with const
package main
import ("fmt")
const FNAME, LNAME = "Josiah", "Olaniyi" 
const AGE =  21
const mSTATUS = "Single"
func main () {
	fmt.Println(FNAME, "\n") // Adding a space vertically
	fmt.Println(LNAME)
	fmt.Println(AGE)
	fmt.Println(mSTATUS)
} */

// Using var
package main
import ("fmt")
var j, o = "Josiah", "Olaniyi"
func main() {
	fmt.Println(j, "", o, "\n") //adding extra space ("")
	fmt.Println(o, "\n") // adding new line
	fmt.Println(j, "\n") // same here
	fmt.Println(o,j) // displayed inversely
}
