package main

import "fmt"

func main(){

	m := make(map[string]int)

	fmt.Println("m: ",m)


	m["one"]=5
	m["two"]=10

	fmt.Println("m :",m)

	v1 := m["one"]
	fmt.Println("v1 : ",v1)


	fmt.Println("length : ",len(m))

	delete(m,"two")
	fmt.Println("After delete : ",m)

	_, prs := m["two"]
	fmt.Println("prs is :",prs)


	n:= map[string]int{"three" :3, "four" : 4}
	fmt.Println("n is :", n)

}
