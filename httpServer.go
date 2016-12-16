/*	Date & Author: Orchid Zhang 20161216
 *	Description: create a HTTP server
 *				temporarily can't work on win10 
 */

package main

import ("fmt"
 "net/http")

func main() {
	
	http.HandleFunc("/", handler)

	http.HandleFunc("/earth", handler2)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Hello Earth\n")
}