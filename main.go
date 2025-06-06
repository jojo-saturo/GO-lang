package main

import (
	"fmt"
)

const FNAME = "Josiah"
const LNAME = "Olaniyi"
const AGE = 21
const ABT = "A young dev developing his skills with GO"

func main() {
	fmt.Print(FNAME, " ", LNAME,"\n")
	fmt.Print(AGE, "\n")
	fmt.Print(ABT)
}
/* Using of Print() doesn't add space and new line automatically, you need add a string for space(" ")
and also you need "\n" for new line */
