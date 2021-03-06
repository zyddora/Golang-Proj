/*	Date & Author: Orchid Zhang 20161216
 *	Description: calculate rectangle and circle area
 *				using Interface 
 */

package main

import ("fmt"
 "math")

func main() {
	
	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))

}

// interface
type Shape interface {

	area() float64
}

type Rectangle struct {

	height float64
	width float64
}

type Circle struct {

	radius float64
}


func (r Rectangle) area() float64{ 
	
	return r.width * r.height
}

func (c Circle) area() float64{ 
	
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64{
	
	return shape.area()
}