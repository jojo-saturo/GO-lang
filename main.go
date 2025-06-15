/*
	// Declearing Mutiples variable:
	package main

	import (

	"fmt"

	) // you need to wrap "fmt" in a bracket

	func main() {
	var a, b = 28, "GO lang"
	c, d := "Var", 30 // while using := you need to asign avariable to it and must be use inside the func main()
	fmt.Println(a, d)
	fmt.Println(c, b)
	}

	package main

	import (
	"fmt"
	)

	var z string = "Yo! what's up?"
	var y int
	var x bool = true

	func main() {
	y = 2025
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	}


	package main
	import "fmt"
	var a, b, c, d = 2, 0, 0, 3 // you can place a varible outside or inside the function
	func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// Display 2003:
	fmt.Println(a, b, c, d)
	}


	// Block Declaration
	package main
	import "fmt"
	func main() {
	var (
	a = "Using Block for declaration"
	b = 100
	c float32 = 9.9
	d = "When using block it must be place inside the function i.e the variable declaration"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d, "\n") //adding "\n" will start a new line in our terminal
	// but is not advicable to use "\n" while printing with .Println

	// you can also print to variables :

	fmt.Println(a, c)
	fmt.Println(b, d)

	// you do not need to add a sting(" ") or \n (space) because println has that as default
	}

	// Using Println()
	package main

	import "fmt"

	var a string
	var b int = 100
	var c bool = true

	func main() {
	a = "Dave"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	}

	// Using Print()
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
	//Using of Print() doesn't add space and new line automatically, you need add a string for space(" ")and also you need "\n" for new line

	//Printf() => deals with argument based on given formatting verb and print them.
	// 2 formatting verbs are:
	// 1.%v: Print value of the argument 2.%T Print type of argument

	package main

	import "fmt"

	func main() {
	var i = "Josiah"
	var h = 50
	fmt.Printf("i has value:%v and type:%T \n", i, i) // adding i,i means we had 2 argument outside the string but inside the bracket
	fmt.Printf("h has value:%v and type:%T", h, h)    // same as this

	// note: using Printf add string inside the bracket and help us make statement.
	}
	package main
	import "fmt"
	func main() {
		var x = true
		var y = 100
		var z = "Hello"
		fmt.Printf("x has value:%v and type:%T \n", x, x)
		fmt.Printf("y has value:%v and type:%T \n", y, y)
		fmt.Printf("z has value:%v and type:%T \n", z, z)
		}
// GO Array
// We declare array using var,[] and type of data we want to store in the array
package main

import (
	"fmt"
)

func main() {
	var arr1 = [4]string{"Josiah", "Olaniyi", "Ade", "Oluwaseun"}
	arr2 := [3]int{3, 2, 1}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// given lenght to your array
	}

package main

import (
	"fmt"
)

func main() {
	var arr1 = [...]string{"Josiah", "Olaniyi", "Bolanle", "Oluwaseun"}
	arr2 := [...]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// without lenght to your array
	}
// changing array object
package main

import (
	"fmt"
)

func main() {
	var cars = [4]string{"Benz", "Mazda", "Nissan", "Volvo"}
	name := [...]string{"Bolanle", "Shalom", "Fortune", "Salomen"}

	cars[3] = "Range-Rover" // changing the array value for cars
	name[0] = "Deborah"     // changing the array value for name

	fmt.Println(cars)
	fmt.Println(name)
	// Printing a spec arr obj
	fmt.Println(cars[2])
	fmt.Println(name[1])
	} */
// Array Initializing
package main

import (
	"fmt"
)

func main() {
	arr1 := [5]int{}              // Not intialized
	arr2 := [5]int{1, 4, 7}       // Partially initailized
	arr3 := [5]int{6, 8, 3, 9, 2} // fully initailized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
