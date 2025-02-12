package main

import "fmt"

	// if you are not using a variables which are declear outside a function then it will not give error
	var name1 string="neeraj"

	var price2=111
	// varience :=10 (you are not able to declear variables with shorthand oprator out side the function)
	

func main() {

	// types of declering variables
	// if you not want to assign values or assign at run time you should declear it only in give (one) format
	var name string = "Bhushan"
	//if variable is not in used it gives error
	
	// without data Type it will assign dynamicaly
	var price =10
	
	// with shortHand Operator
	varience :=10.1003

	fmt.Println(name)
	fmt.Println(price)
	fmt.Println(varience)
}