package main

type animal struct{
	name string `required max:"100"`
	origin string
}

type bird struct{
	animal //similar to inherit
	speed float32
	canFly bool
}