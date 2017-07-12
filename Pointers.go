package main

import "fmt"

func zeroval(i int) {

	i = 10

}

func zerovalptr(iptr *int) {

	*iptr = 10

}


func main()  {

	val := 3

	fmt.Println("val is :",val)

	zeroval(3)
	fmt.Println("val of i is :",val)

	zerovalptr(&val)

	fmt.Println("pointer value",val)

	fmt.Println("pointer",&val)
}


