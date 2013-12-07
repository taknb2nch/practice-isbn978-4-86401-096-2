package main

import (
	"fmt"
	"math"
)

type cartesianPont struct {
	x, y float64
}

type polarPoint struct {
	r, θ float64
}

func (p cartesianPont) X() float64 {
	return p.x
}

func (p cartesianPont) Y() float64 {
	return p.y
}

func (p polarPoint) X() float64 {
	return p.r * math.Cos(p.θ)
}

func (p polarPoint) Y() float64 {
	return p.r * math.Sin(p.θ)
}

func (self cartesianPont) Print() {
	fmt.Printf("(%f, %f)\n", self.x, self.y)
}

func (self polarPoint) Print() {
	fmt.Printf("(%f, %f度\n", self.r, self.θ)
}

type Point interface {
	Printer
	X() float64
	Y() float64
}

// メソッドが1つの場合はメソッド+er
type Printer interface {
	Print()
}

func MakeCartesianPoint(x, y float64) Point {
	return cartesianPont{x, y}
}

func MakePolarPoint(x, θ float64) Point {
	return polarPoint{x, θ}
}

func main() {
	// var p Printer
	// point := MakeCartesianPoint(1, 2)
	// p = point

	p := MakeCartesianPoint(1, 2)

	p.Print()

	p = MakePolarPoint(5, 45)

	p.Print()
}