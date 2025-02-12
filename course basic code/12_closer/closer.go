package main

import "fmt"

func me() func() int {
	count := 0
	return func() int {
		count=count+1
		return count
	}
}

func main() {
	increment := me()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}