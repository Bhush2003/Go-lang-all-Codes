package main
import (
	"fmt"
)

type OrderStatus int

// also give string
const(
	Recived OrderStatus =iota
	Notrecived
	Prepared
) 

func order(status OrderStatus){
	fmt.Println(status)
}

func main(){
	order(Prepared)
	order(Notrecived)
}
