package main

import "fmt"

type react struct {

	width, height int
}

func (r react) area() int {

	return r.width * r.height
}

func (r react) perimi() int {

	return 2*r.width + 2*r.height
}


func main()  {

	r := react{width:10,height:5}

	fmt.Println(r)


	fmt.Println("Area is : ",r.area())
	fmt.Println("Permi is : ",r.perimi())
}
