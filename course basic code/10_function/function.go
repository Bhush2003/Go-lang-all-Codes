package main

import "fmt"


func me1() (string, int, bool){
	return "1111",10,true
} 
func me() int{
	return 10
} 

func main(){

	me()
	// s1,s2,_:=me1()
	s1,s2,s3:=me1()
	fmt.Println(s1,s2,s3)
}