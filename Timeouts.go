package main

import (
	"time"
	"fmt"
)

func main()  {

	c1 := make(chan string,1)

	go func() {

		time.Sleep(time.Second *1)
		c1 <-"one"

	}()

	fmt.Println("channel : ",<-c1)
}
