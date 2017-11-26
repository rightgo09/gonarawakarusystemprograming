package main

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//main3_9_1_1()
	//main3_9_1_2()
	//main3_9_1_3()
	//main3_9_1_4()
	//main3_9_1_5()
	main3_9_1_6()
}

func main3_9_1_1() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)
}

func main3_9_1_2() {
	file, err := os.Create("rand.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := make([]byte, 1024)
	//var buffer bytes.Buffer
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	reader := bytes.NewReader(b)
	//io.CopyN(file, reader, 1024)
	io.Copy(file, reader)
}

func main3_9_1_3() {
	file, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("archive.txt")
	if err != nil {
		panic(err)
	}
	writer.Write([]byte("FOO BAR BAZ"))
}

func main3_9_1_4() {
	http.HandleFunc("/", main3_9_1_4_handler)
	http.ListenAndServe(":8080", nil)
}

func main3_9_1_4_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii/sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("foo.txt")
	if err != nil {
		panic(err)
	}
	writer.Write([]byte("YES NO PRICURE!"))
}

func main3_9_1_5() {
	source := "abcdefghijklmn"
	myiocopyn(os.Stdout, strings.NewReader(source), 400)
}

func myiocopyn(dest io.Writer, src io.Reader, length int) {
	//io.Copy(dest, io.NewSectionReader(src, 0, length))
	io.Copy(dest, io.LimitReader(src, int64(length)))
}

func main3_9_1_6() {
	computer := strings.NewReader("COMPUTER")
	system := strings.NewReader("SYSTEM")
	programming := strings.NewReader("PROGRAMMING")

	var stream io.Reader

	// ここにioパッケージを使ったコードを書く
	a := io.NewSectionReader(programming, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i := io.NewSectionReader(programming, 8, 1)
	//i2 := io.NewSectionReader(programming, 8, 1)
	//
	//stream = io.MultiReader(a, s, c, i, i2)
	pr, pw := io.Pipe()
	writer := io.MultiWriter(pw, pw)
	go io.CopyN(writer, i, 1)
	defer pw.Close()
	stream = io.MultiReader(a, s, c, io.LimitReader(pr, 2))
	//stream = io.MultiReader(a, s, c, pr)

	io.Copy(os.Stdout, stream)
}
