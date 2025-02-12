package main

import (
	"fmt"
	"time"
)

func main(){

	i:=1

	// similer like java type of switch statements
	switch i{
	case 1:
		fmt.Println("printing")
	case 2:
		fmt.Println("not printing")
	default:
		fmt.Println("else")
	}

	// multiple condition in cases
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("not printing")
	case time.Thursday,time.Monday,time.Tuesday,time.Friday,time.Wednesday:
		fmt.Println("printing")
	default:
		fmt.Println("ather wise printing")
	}
}

