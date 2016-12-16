/*	Date & Author: Orchid Zhang 20161216
 *	Description: go routine, parallel
 *				1. the example video runs out a 
 *				consecutive number executed.
 *				2. but when executed on Win10, the 
 *				sequence of output is changed every execution.
 *				3. Personally think that kind of optimization 
 *				had been done with all the routines.
 */

package main

import ("fmt"
 "time")

func count(id int) {
	
	for i := 0; i < 10; i++{
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	
	for i := 0; i < 10; i++{
		go count(i)
	} 

	time.Sleep(time.Millisecond * 11000)
}
