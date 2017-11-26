package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	main3_7_1_2()
}

func main3_7_1_1() {
	header := bytes.NewBufferString("------------ HEADER -------------\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("------------ FOOTER -------------\n")
	reader := io.MultiReader(header, content, footer)
	// すべてのreaderをつなげた出力が表示
	io.Copy(os.Stdout, reader)
}

func main3_7_1_2() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReaer\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// けどバッファには残ってる
	fmt.Println(buffer.String())
}
