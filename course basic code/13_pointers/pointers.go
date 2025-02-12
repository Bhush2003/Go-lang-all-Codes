package main

import "fmt"

func me(i *int) {
	*i = 5
}

func main() {
	num := 1
	me(&num)
	fmt.Println(num)
}