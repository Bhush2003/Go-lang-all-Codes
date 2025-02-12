package main

import "fmt"



// keyword for generics (any, interface{}, comparable)
func multipalconparable[T int|string|bool](slic []T,name T){
	for _ ,i:=range slic{
		fmt.Println(i,name)
	}
}
func multipalconparable1[T comparable](slic []T,name T){
	for _ ,i:=range slic{
		fmt.Println(i,name)
	}
}


func main(){
	slic:=[]int{1,2,3,4,5}
	multipalconparable(slic,3)
	multipalconparable1(slic,5)
}