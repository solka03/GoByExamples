package main

import "fmt"

type person struct{

	name string
	age int

}


func main()  {

	fmt.Println(person{name : "alan",age :16})
	fmt.Println(person{name:"rayn"})
	fmt.Println(person{age:22})

	sp := person{name:"twin",age:22}

	fmt.Println(sp)

	val := &sp

	fmt.Println("age is ",val.age)

}
