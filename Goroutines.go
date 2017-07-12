package main

import (
	"fmt"
	"time"
)

func func1(from string)  {

	for i:=0;i<len(from);i++ {

		fmt.Println(from," : ",i)
	}
}

func main()  {

	func1("direct")

	time.Sleep(30000)
	go func1("goroutine")

	go func(message string){
		fmt.Println(message)
	}("going")

	var input string
	fmt.Println(&input)
	fmt.Println("done")

}
