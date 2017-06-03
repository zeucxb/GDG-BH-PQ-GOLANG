package main

import (
	"GDG-BH-PQ-GOLANG/tasks"
	"fmt"
)

func main() {
	done := make(chan int)

	go tasks.HardTask(done)
	tasks.Task()
	tasks.Task()
	tasks.Task()

	for {
		select {
		case r := <-done:
			fmt.Println("Soma =", r)
			return
		default:
			fmt.Println(".")
		}
	}

}
