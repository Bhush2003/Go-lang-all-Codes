package main

import "fmt"

func main() {
	age := 10
	if age >= 18 {
		fmt.Println("person is adult")
	}else if age<=12{
		fmt.Println("person is child")
	}else{
		fmt.Println("pesrson is teen")
	}
	// you can use && || oprators like java in go
	// go does not have tunary oprator
	// you can decleare a variable in condition in flutter

	if age := 12; age<13{
		fmt.Println("bhushan")
	}

}