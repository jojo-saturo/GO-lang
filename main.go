/*
// Declearing varibles
package main
import ("fmt")
var a int
var b int = 2
var c int = 3
func main () {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Declearing multiple variables
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 2, 0, 0, 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// Declearing multiple with int and stirng
package main

import (
	"fmt"
)

func main() {
	var a, b = 2, "Jojo"
	c, d := 8, "Saturo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
} */

// Declearation in Block
package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b string  = "Jojo saturo BD"
		c bool    = true
		d float32 = 28.06
	)
	a = 2003
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
