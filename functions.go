package main

import "fmt"

func main()  {

	res := plus(2,3)
	fmt.Println("res : ", res)

	resp := plusplus(4,5,6)
	fmt.Println("resp : ", resp)
}

func plus( a int, b int)  int{

	return a+b
}

func plusplus(a,b,c int) int {

	return a+b+c
}