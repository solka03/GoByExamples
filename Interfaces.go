package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area1() float64
	perimiti() float64
}

type rect struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area1() float64{

	return float64(r.width * r.height)
}

func (r rect) perimiti() float64{

	return float64(2*r.width + 2*r.height)
}

func (c circle) area1() float64 {

	return float64( math.Pi * c.radius * c.radius)
}

func (c circle) perimiti() float64  {
	return float64(2 * math.Pi * c.radius)
}

func measure(g geometry)  {
	fmt.Println("Area :",g.area1())
	fmt.Println("Perimiti : ",g.perimiti())
}

func main()  {

	r := rect{10,5}
	c := circle{6}

	fmt.Println("Rectangle :",r)
	fmt.Println("Circle :",c)


	measure(r)
	measure(c)



}
