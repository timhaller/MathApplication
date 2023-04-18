package main

import (
	"example.com/Newton"
	"fmt"
	"time"
)

func main() {

	startTime := time.Now().UnixMilli()

	xI := float64(1)

	for i := 0; i < 10; i++ {
		xI = Newton.X(xI)
		fmt.Println(i, ":", xI)
	}

	endTime := time.Now().UnixMilli()

	fmt.Println("This process took", endTime-startTime, "ms to run")

}
