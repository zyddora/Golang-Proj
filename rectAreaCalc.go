/*	Date & Author: Orchid Zhang 20161216
 *	Description: Golang struct, pointer, function 
 */

package main

import "fmt"

func main() {
	
	rect1 := Rectangle{0, 50, 10, 10}

	fmt.Println("Area of rectangle =", rect1.area())

}

type Rectangle struct {

	leftX float64
	topY float64
	height float64
	width float64

}

// focus on the input para with pointer
func (rect *Rectangle) area() float64{ 
	
	return rect.width * rect.height

}