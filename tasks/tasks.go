package tasks

import (
	"fmt"
	"time"
)

func Task() {
	fmt.Println("INIT Task")
	time.Sleep(time.Second * 1)
	fmt.Println("END Task")
}

func HardTask(done chan int) {
	fmt.Println("INIT Hard Task")
	time.Sleep(time.Second * 5)
	fmt.Println("END Hard Task")
	done <- 9 + 189
}
