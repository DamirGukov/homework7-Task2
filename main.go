package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	arr := make([]int, 0, 3)

	rand.Seed(time.Now().UnixNano())

	go func() {
		for i := 0; i <= 3; i++ {
			n := rand.Intn(100)
			fmt.Println("Random number:", n)
			c <- n
		}
		close(c)
		
		answerMax := <-c2
		fmt.Println("Max number:", answerMax)
		answerMin := <-c3
		fmt.Println("Min number:", answerMin)

	}()
	time.Sleep(time.Second)

	go func() {
		for num := range c {
		arr = append(arr, num)
	}
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

