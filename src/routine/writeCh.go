package routine

import (
	"fmt"
	"time"
)

func writeToChannel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func WriteCH() {
	c := make(chan int)
	go writeToChannel(c, 10)

	time.Sleep(1 * time.Second)
	fmt.Println("Read", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is Open!")
	} else {
		fmt.Println("Channel is Closed!")
	}
}
