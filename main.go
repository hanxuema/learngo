package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	n:= 42
	fmt.Printf("%v, %T", n,n)
	
	s := "this is string"
	fmt.Println(s)
	b := []byte(s)
	fmt.Printf("%v,%T\n",b,b)
	
	grades := [3]int{97,85,93}
	fmt.Printf("Grades : %v", grades)
}
