package main

import (
	"fmt"
	"math"
)

const s  = "string constant" 
func main()  {

	const n  = 5000000
	fmt.Println(n/2)

	const d  = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
