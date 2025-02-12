package main

import "fmt"

const age=30
const name = "bhushan"
func main(){

	const(
		port = 200
		host = "me"
	)
	// with "," it will print both variables
	fmt.Println(age,name)
	fmt.Println(port,host)

}