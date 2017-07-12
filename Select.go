package main

import (
	"time"
	"fmt"
)

func main()  {

	c1 := make(chan string,1)
	c2 := make(chan string,1)

	go func() {

		time.Sleep(time.Second *1)
		c1 <- "one"
	}()


	go func() {

		time.Sleep(time.Second *1)
		c2 <- "two"

	}()

	for i:=0;i<2;i++ {
		select {
		case msg := <-c1:
			fmt.Println("c1 channel :",msg)
		case msg :=<-c2:
			fmt.Println("c2 channel",msg)
		}
	}

}