package main

import "fmt"

func main()  {

	two := sum(12,22)
	fmt.Println("two result :",two)

	three := sum(2,4,5)
	fmt.Println("three result : ",three)



}

func sum(nums ...int) int {

	fmt.Println("numbers is ", nums)

	total := 0

	for _,num := range nums {

		total +=num
	}
	return total
}

