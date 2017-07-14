package main

import (
	"os"
)

func main() {

	panic("there is a problem")

	_,err := os.Create("/tmp/save")
	if err != nil{
		panic(err)
	}

}
