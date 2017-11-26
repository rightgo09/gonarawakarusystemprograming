package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	//main2_4_6_1()
	main2_4_6_2()
}

func main2_4_6_1() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}

func main2_4_6_2() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
