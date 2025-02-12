package main
import "fmt"

// in go there is no while key word
// there is only for keyword

func main(){
	// while loop in go with for keyword
	i :=0
	for i<=3{
		fmt.Println(i)
		i=i+1
	}

	fmt.Println()

	// for loop implimentation
	for j:=0;j<=3;j++{
		fmt.Println(j)
	}

	fmt.Println()

	// range keyword like python in go
	for j:=range 4{
		fmt.Println(j)
	}
}