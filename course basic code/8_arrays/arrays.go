package main

import "fmt"

func main(){
	// if no element is assign then zerod values are assign 
	var name [3]int
	name[0]=1
	fmt.Println(name[0])

	var name1 [4]string
	name1[0]="bhushan"
	fmt.Println(name1)

	// another type of declering array

	name2 :=[5]bool{}

	fmt.Println(name2)


	
}