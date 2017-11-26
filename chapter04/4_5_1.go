package main

import (
	"time"
	"fmt"
)

func main() {
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("10 byou tattayo!")
			return
		}
	}
}
