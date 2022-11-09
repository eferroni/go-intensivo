package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

// T1 -> main
func main() {
	// go task("A") // go routine // green threads
	// go task("B")
	// task("C")

	canal := make(chan string)

	// T2
	go func() {
		canal <- "Veio da t2"
	}()

	// T1
	fmt.Println(<-canal)
}