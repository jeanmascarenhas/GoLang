package main

import (
	"fmt"
)

const x = 10

const (
	a = 10
	b = 20
	e = 30
)

//var y = 10

var c int
var d float64

func Consts() {
	c = x
	d = x

	fmt.Println("value", c, d)
	fmt.Println("Consts", a, b, e)
}

func BitsDislocation() {
	x := 24
	y := x << 2
	z := x >> 2

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
