package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, playground")
	// n:= 42
	// fmt.Printf("%v, %T", n,n)
	
	// s := "this is string"
	// fmt.Println(s)
	// b := []byte(s)
	// fmt.Printf("%v,%T\n",b,b)
	
	// grades := [3]int{97,85,93}
	// fmt.Printf("Grades : %v", grades)

	a := [...]int{1,2,3}
	b := &a // make b to be the pointer to a
	b[1] = 5

	c := a //create another array of c
	c[1] =7
	fmt.Println("c is ",c)
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
	fmt.Printf("Length: %v\n", len(a))  //%v default value
	fmt.Printf("Capacity: %v\n", cap(a))  //%v default value
	
	a1 := []int{1,2,3,4,5,6,6,7,8,9,10}
	s1 := a1[:] //slice all elements
	s2 := a1[3:] //slice from 4th to end
	s3 := a1[:6] //slice first 6 elemetns
	s4 := a1[3:6] //slice from 4th, 5th and 6th

	fmt.Println("a1 is ", a1)
	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)
	fmt.Println("s3 is ", s3)
	fmt.Println("s4 is ", s4)
}
