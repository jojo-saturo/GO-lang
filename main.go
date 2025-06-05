package main

import "fmt"

var studentName string
var studentNum int
var studentGrade float32
var student bool

func main() {
	studentName = "Jojo Saturo"
	studentNum = 500
	studentGrade = 4.5
	student = true

	fmt.Println(studentName)
	fmt.Println(studentNum)
	fmt.Print(studentGrade)
	fmt.Println(student)
}
