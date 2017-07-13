package main

import (
	"sort"
	"fmt"
)

func main()  {

	strs := []string{"c","b","d"}
	sort.Strings(strs)
	fmt.Println("string sorted : ",strs)

	ints := []int{4,9,8,7,6,10}
	sort.Ints(ints)
	fmt.Println("integer sorted : ",ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted : ",s)

}