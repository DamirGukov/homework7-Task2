package main

import (
	"fmt"
	"math/rand"
	"time"
)

type minMax struct {
	min int
	max int
}

func main() {
	c := make(chan int)
	c2 := make(chan minMax)

	arr := make([]int, 0, 3)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for i := 0; i < 3; i++ {
			n := rand.Intn(100)
			fmt.Println("Random number:", n)
			c <- n
		}
		close(c)

		answer := <-c2
		fmt.Println("Min number:", answer.min)
		fmt.Println("Max number:", answer.max)
	}()

	time.Sleep(time.Second)

	go func() {
		for num := range c {
			arr = append(arr, num)
		}

		maxMin := minMax{
			min: arr[0],
			max: arr[0],
		}

		for _, element := range arr {
			if element > maxMin.max {
				maxMin.max = element
			}
			if element < maxMin.min {
				maxMin.min = element
			}
		}

		c2 <- maxMin

	}()

	time.Sleep(time.Second)
}

