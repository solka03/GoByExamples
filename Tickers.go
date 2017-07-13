package main

import (
	"time"
	"fmt"
)

func main()  {

	tickers := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range tickers.C {
			fmt.Println("ticker at : ",t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)

	tickers.Stop()

	fmt.Println("stopped ticker")


	}

