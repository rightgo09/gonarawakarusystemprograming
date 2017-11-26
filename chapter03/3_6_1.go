package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行め
2行め
3行め`

func main() {
	main3_6_1_2()
}

func main3_6_1_1() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

func main3_6_1_2() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
