package main

import "fmt"

func sum(nums ...int)int{
	total:=0
	for _, num:=range nums{
		total=total+num
	}
	
	return total
}

func sum1(nums ...interface{}){
	for _, num:=range nums{
		fmt.Print(num)
	}
}

func main(){
	m:=sum(1,2,3,4,4,5,6)
	
	fmt.Println(m)
	sum1(1,2,3,4,4,5,6,7,8,9,10)

}