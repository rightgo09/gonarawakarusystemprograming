package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"time"
)

func main() {
	//main2_4_5_1()
	//main2_4_5_2()
	main2_4_5_3()
}

func main2_4_5_1() {
	file, err := os.Create("multiwritter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.Multiwriter example\n")
}

func main2_4_5_2() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}

func main2_4_5_3() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	time.Sleep(1 * time.Second)
	buffer.WriteString("example\n")
	buffer.Flush()
}
