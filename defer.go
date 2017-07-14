package main

import (
	"os"
	"fmt"
)

func main()  {

	f := createFile("/tmp/test.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(fileName string) *os.File {
	fmt.Println("Creating File")
	f,err := os.Create(fileName)
	if err != nil{
		panic(err)
	}
	return f

}

func writeFile(f *os.File)  {
	fmt.Println("writing into File")
	fmt.Fprintln(f,"writing")
}

func closeFile(f *os.File)  {
	fmt.Println("close the file")
	f.Close()
}
