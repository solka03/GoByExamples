package main

import "fmt"

func intseq() func() int{

	i := 0

	fmt.Println("i : ==",i)
	return func() int {

		i += 1
		return i
	}

}

func main()  {


	nextInt := intseq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt1 := intseq()
	fmt.Println(nextInt1())


}

