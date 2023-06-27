package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c2 := make(chan int)
	c3 := make(chan int)
	arr := make([]int, 0, 3)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for i := 0; i < 3; i++ {
			n := rand.Intn(100)
			fmt.Println("Random number:", n)
			arr = append(arr, n)
		}
		answerMax := <-c2
		fmt.Println("Max number:", answerMax)
		answerMin := <-c3
		fmt.Println("Min number:", answerMin)

	}()

	time.Sleep(time.Second)

	go func() {
		max := arr[0]
		min := arr[0]
		for _, element := range arr {
			if element > max {
				max = element
			}
			if element < min {
				min = element
			}
		}
		c2 <- max
		c3 <- min

	}()

	time.Sleep(time.Second)
}
