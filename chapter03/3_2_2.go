package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("abcdefghijklmn")
	io.CopyN(os.Stdout, r, 4)
	fmt.Println("")

	buffer := make([]byte, 4)
	io.CopyBuffer(os.Stdout, r, buffer)
	fmt.Println("")
}
