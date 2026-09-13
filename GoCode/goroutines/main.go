package main

import (
	"fmt"
	"time"
)

func main() {
	go dosomething()
	fmt.Printf("line in main after fn with for loop\n")
	for {
		time.Sleep(7 * time.Second)
		break
	}

}

func dosomething() {
	fmt.Printf("first line in fn before sleep\n")
	time.Sleep(5 * time.Second)
	fmt.Printf("last line in fn after sleep\n")
}
