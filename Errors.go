package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int,error)  {

	if arg == 42{

		return -1,errors.New("42 is wrong one")
	}

	return arg+3,nil
}

type argError struct {
	arg int
	probe string
	}

func (r *argError) Error() string {
	return fmt.Sprintf("%d - %s", r.arg, r.probe)
}

func f2(arg int) (int,error)  {

	if arg == 42{
		return  -1,&argError{arg,"can't work with 42"}
		
	}

	return arg+3,nil

}

func main()  {

	for _,i := range []int{7,42}{

		if v,e := f1(i); e != nil{
			fmt.Println("error in f1",e)
		}else{
			fmt.Println("f1 worked",v)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.probe)
	}


}