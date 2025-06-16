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
	}
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
// Initailizing Specific Element
package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{5: 70, 2: 50, 0: 55}

	fmt.Println(arr1)
	fmt.Println(arr1[0])

	// 5: 70 means assign 70 to 5th (4)element
	// 2: 50 means assign 50 to 2nd (3)element
	// 0: 55 means assign 55 to zero(1)element
	}

// Printing lenght of an Array
package main

import (
	"fmt"
)

func main() {
	arr1 := [...]string{"Monitor", "Gate", "Master-class", "Book", "Island"}
	arr2 := [...]int{1, 4, 5, 7, 8, 33, 45, 70, 3, 22}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}

// GO Slices
// Creating slice with;
package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}

	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	// A slice with 0 len and cap

	myslice2 := []string{"GO", "Slices", "Are", "Powerful"}

	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// A slice with 4 len and cap
	}

// Creating Slice from an Array
// var myarray = [lenght]datatype{value}
// myslice := myarray [start:end]

package main

import (
	"fmt"
)

func main() {
	var arr = [6]int{3, 6, 9, 5, 7, 1}
	myslice := arr[1:4]

	fmt.Printf("myslice = %v\n", myslice)
	// Printing the value of the arr

	fmt.Printf("lenght = %d\n", len(myslice))
	// Pirnting the lenght of the array from above

	fmt.Printf("Capacity = %d\n", cap(myslice))
	// The capacity is 5 because the slice start from 1
	}

// Crearing a Slice with make() function
// slice_name := make([]type, lenght, capacity)

package main

import (
	"fmt"
)

func main() {
	// Assigning Capacity

	myslice1 := make([]int, 5, 10)

	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("lenght = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// Without Assigning Capacity

	myslice2 := make([]int, 7)

	fmt.Printf("myslice = %v\n", myslice2)
	fmt.Printf("lenght = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	// The capacity is set to the lenght number.
	}

// Changing a slice element

package main

import (
	"fmt"
)

func main() {
	myslice := []int{3, 5, 6, 4, 9}
	myslice[2] = 7
	myslice[3] = 9
	myslice[4] = 11

	fmt.Printf("myslice 3 (6 => 7) = %v\n", myslice[2])
	fmt.Printf("lenght = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	fmt.Printf("myslice = %v\n", myslice)
	}

// Apend Element to a Slice
// slice_name = append(slice_name, element1, element2, ...)

package main

import (
	"fmt"
)

func main() {
	myslice1 := []string{"Book", "Maximize", "Root", "Domain"}
	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("lenght = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// Appending to the the element

	myslice1 = append(myslice1, "Loyalty", "Universe")
	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("lenght = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	}*/

//Apending One Slice to another
// slice_name := append(slice1, slice2...)

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{2, 4, 6}
	myslice2 := []int{3, 5, 7}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("lenght = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
}
