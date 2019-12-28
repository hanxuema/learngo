package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

	a := [...]int{1, 2, 3}
	b := &a // make b to be the pointer to a
	b[1] = 5

	c := a //create another array of c
	c[1] = 7
	fmt.Println("c is ", c)
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
	fmt.Printf("Length: %v\n", len(a))   //%v default value
	fmt.Printf("Capacity: %v\n", cap(a)) //%v default value

	a1 := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}
	s1 := a1[:]   //slice all elements
	s2 := a1[3:]  //slice from 4th to end
	s3 := a1[:6]  //slice first 6 elemetns
	s4 := a1[3:6] //slice from 4th, 5th and 6th first item is inclusive and second item is exclusive

	fmt.Println("a1 is ", a1)
	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)
	fmt.Println("s3 is ", s3)
	fmt.Println("s4 is ", s4)

	a3 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("a3 is ", a3)
	fmt.Println("a3[:2] is ", a3[:2])
	fmt.Println("a3[3:]... is ", a3[3:])
	b2 := append(a3[:2], a3[3:]...)
	fmt.Println("b2 is ", b2)
	fmt.Println("a3 is ", a3)

	s5 := make([]int, 10)      //create slice with capacity and length = 10
	s6 := make([]int, 10, 100) //create slice with length = 10 and capacity = 100

	fmt.Println("s5 is ", s5)
	fmt.Println("s6 is ", s6)

	dic := make(map[string]int)
	dic = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(dic)
	dic["a"] = 5
	fmt.Println(dic)
	delete(dic, "a") //delete a
	fmt.Println(dic["a"])
	fmt.Println(dic) // a is zero even a is deleted

	pop, ok := dic["a"] //fix the above issue
	fmt.Println(pop, ok)

	fmt.Println(len(dic))

	aDoctor := struct{ name string }{name: "xavier"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom"
	fmt.Println("aDoctor is", aDoctor)
	fmt.Println("anotherDoctor is ", anotherDoctor)
	aDoctor.name = "abb"
	fmt.Println("aDoctor is", aDoctor) //because struct type is value type, therefore anotherdoctor 's name is not affacted by adoctor
	fmt.Println("anotherDoctor is ", anotherDoctor)

	//use pointer
	secondDoctor := &aDoctor
	fmt.Println("aDoctor is", aDoctor)
	fmt.Println("second doctor is", secondDoctor)
	aDoctor.name = "jlsadfjlasdf"

	fmt.Println("aDoctor is", aDoctor)
	fmt.Println("second doctor is", secondDoctor)
	fmt.Println("anotherDoctor is ", anotherDoctor)

	bird1 := bird{}
	bird1.name = "emu"
	bird1.origin = "australia"
	bird1.speed = 87
	bird1.canFly = false

	fmt.Println(bird1)

	t := reflect.TypeOf(animal{})
	fmt.Println(t)
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

	//below will run start and end and middle
	fmt.Println("start")
	//defer fmt.Println("middle") // defer get executed after main executes but before it is returnd
	fmt.Println("end")

	//defer to last in first out, stack
	//below will be end, middle, start

	// defer fmt.Println("start 1")
	// defer fmt.Println("middle 2")
	// defer fmt.Println("end 3")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = robots
	// fmt.Printf("%s", robots)

	//it will only print deferA, because deferA is used by defer function, deferA will only be delcared when the deferFunc get called
	// deferA := "deferA"
	// defer fmt.Println(deferA)
	// deferA = "deferBBBBBBBBBB"

	// fmt.Println("start")

	// defer func()  {
	// 	if err:= recover(); err != nil {
	// 		log.Println("error", err)
	// 	}
	// }()
	// panic("something bad happend")
	// fmt.Println("end")

	aaa := 42
	bbb := aaa
	bbbPointer := &aaa // bbbPointer is ther pointer to aaa
	fmt.Println(aaa, bbb, bbbPointer)
	aaa = 76
	//get the value of pointer, which is the value of a, it output 76 42 76
	fmt.Println(aaa, bbb, *bbbPointer)

	aa1 := [3]int{1, 2, 3}
	bb1 := &aa1[0]
	cc1 := &aa1[1]
	fmt.Printf("%v %p %p\n", aa1, bb1, cc1) //print the value of a and the pointer of b and c

	var ms *myStruct
	fmt.Println(ms) //<nil>
	ms = new(myStruct)
	fmt.Println(ms) //&{0}
	ms.foo = 78
	fmt.Println(ms.foo) //78
	(*ms).foo = 42
	fmt.Println((*ms).foo) //42

	greeting := "hello"
	name := "Stan"
	sayGreeting(greeting, name)
	fmt.Println(name) //print stan

	sayGreetingPointer(&greeting, &name)
	fmt.Println(name) //print ted

	println("outside", sum(1, 2, 3, 4, 5))

	div, err := divide(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(div)

	div, err = divide(1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(div)

	for index := 0; index < 5; index++ {
		func(i int) {
			fmt.Println("Hello go", i)
		}(index)
	}
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name) //hello stan
	name = "Ted"                // from stan to ted
	fmt.Println(name)           //print ted
}

func sayGreetingPointer(greeting, name *string) {
	fmt.Println(greeting, name)   //0xc00008cfe0 0xc00008cff0
	fmt.Println(*greeting, *name) //hello stan
	*name = "Ted"                 // from stan to ted
	fmt.Println(name)             //print 0xc00008cff0
}

//... must be the last one and the only one parameter
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("sum is", result)
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}
	return a / b, nil
}

type myStruct struct {
	foo int
}
