/*	Date & Author: Orchid Zhang 20161216
 *	Description: parallel example - make pizza
 */

package main

import ("fmt"
 "time"
 "strconv")

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and send for sause")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSause(stringChan chan string) {
	
	pizza := <- stringChan

	fmt.Println("Add Sause and Send", pizza, "for toopings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	
	pizza := <- stringChan

	fmt.Println("Add Toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)
}

func main() {
	
	stringChan := make(chan string)

	for i := 0; i < 3; i++{

		go makeDough(stringChan)
		go addSause(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)

		fmt.Println("---------------------")
	}
}
