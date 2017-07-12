package main

import "fmt"

func main()  {

	var a [5]int

	for i:=0;i<5;i++ {

		fmt.Println("a[i] :",i,a[i])
	}

	a[4] = 100

	fmt.Println("a[4] : ",a[4])

	b := [5]int {2,4,6,7,8}

	fmt.Println("b[i] :",b)

}
