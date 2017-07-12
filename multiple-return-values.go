package main

import "fmt"

func main(){

	a,b := multi()

	fmt.Println("a is :", a)
	fmt.Println("b is :",b)

	_,c := multi();
	fmt.Println("c is :",c)

}

func multi() (int,int){

	return 3,4
}
