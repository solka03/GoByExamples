package main

import "fmt"

func main()  {

	jobs := make(chan int,5)
	done := make(chan bool)

	go func() {

		for {

			j,more := <-jobs
			if more {
				fmt.Println("Received job ", j)
			}else{
				fmt.Println("Received All jobs")
				done <- true
				return
			}
			
		}

	}()

	for i:=0; i < 3;i++ {
		jobs <- i
		fmt.Println("sent job  ", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-done




}