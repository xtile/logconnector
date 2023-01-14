package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main function
func main() {

	rand.Seed(42)

	i := 0

	for {

		i++

		time.Sleep(1000 * time.Millisecond)

		//fmt.Printf("%v Test log record \n", rand.Intn(10000))
		fmt.Printf("%v Test log record \n", i)

	}

}
