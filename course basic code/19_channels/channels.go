package main

import (
	"fmt"
	"time"
)



func me[T int|string] (chan1 chan T){
    fmt.Println("num is", <-chan1)
}

func main(){
    chan1:=make(chan int)
    chan2:=make(chan string)
    
    go me(chan1)
    go me(chan2)
    chan1 <- 5
    chan2 <- "bhushan"

    time.Sleep(time.Second)
}


