package main

import "fmt"

func ping(pings chan<- string, msg string)  {

	pings <- msg

}

func pong(pongs chan<- string,pings <-chan string)  {
	msg := <-pings
	pongs <- msg

}


func main()  {

	pings := make(chan string,1)
	pongs := make(chan string,2)

	ping(pings,"passed message")
	pong(pongs,pings)

	fmt.Println(<-pongs)
}