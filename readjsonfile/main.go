package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	var filename string
	flag.StringVar(&filename, "f", "test.json", "sample json file")
	flag.Parse()

	if filename == "" {
		return
	}

	file, err := os.Open(filename)

	if err != nil {
		return
	}

	byteArr, err := io.ReadAll(file) // ioutil.ReadAll(file)
	if err != nil {
		return
	}

	var jsonMap map[string]interface{}

	err = json.Unmarshal(byteArr, &jsonMap)
	if err != nil {
		return
	}

	for k, v := range jsonMap {
		fmt.Printf("for key %s the value is %v \n", k, v)
	}
}
