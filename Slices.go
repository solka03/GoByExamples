package main

import "fmt"

func  main()  {

	s:=make([]string,3)
	fmt.Println("emp :",s)

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"

	fmt.Println("s :", s)
	fmt.Println("s[2]",s[2])

	c := make([]string, len(s))

	copy(c,s)

	fmt.Println("c :", c)
	fmt.Println("c[2] :",c[2])




}
