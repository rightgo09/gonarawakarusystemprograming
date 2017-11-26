package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	first := make(chan string)
	second := make(chan string)

	i := 0

	go func() {
		time.Sleep(3 * time.Second)
		second <- "second no channel kara kitayo!"
	}()

	go func() {
		A:
		for {
			select {
			case d1 := <-first:
				fmt.Println(d1)
			case d2 := <-second:
				fmt.Println(d2)
				break
			default:
				i += 1
				fmt.Println("default", i)
				if i > 100 {
					fmt.Println("break!")
					break A
				}
			}
		}

		done <- struct{}{}
	}()

	first <- "first no channel kara kitayo!"
	first <- "first no channel kara kitayo!"
	first <- "first no channel kara kitayo!"
	first <- "first no channel kara kitayo!"
	first <- "first no channel kara kitayo!"

	<-done
}
