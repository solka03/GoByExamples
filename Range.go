package main

import "fmt"

func main()  {

	nums := []int{2,3,4}

	sum := 0
	for _,num := range nums{

		sum += num

	}

	fmt.Println("total is ;",sum)


	for i,three := range nums {

		if three == 3{
			fmt.Println("index of 3 is :",i)
		}

	}

	kvp := map[string]string{ "a":"apple", "b":"banana" }

	for k,v := range kvp{
		fmt.Println("key :",k, " value:",v)
	}

	for key := range kvp {

		fmt.Println("key", key)
	}

	for i,c := range "go"{

		fmt.Println("i :",i," c:",c)

	}


}
