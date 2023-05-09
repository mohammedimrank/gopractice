package main

import (
	"fmt"
	"net/http"
	"strings"
)

var userMap map[string]int

func main() {

	userMap = make(map[string]int)
	///GET /user/{name}
	http.HandleFunc("/user/", getUserByName)
	http.ListenAndServe("0.0.0.0:5000", nil)

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUserByName(resp http.ResponseWriter, req *http.Request) {

	username := strings.TrimPrefix(req.URL.Path, "/user/")
	var count = 1

	if c, ok := userMap[username]; ok {
		count = c
	}

	userMap[username] = count + 1
	resval := fmt.Sprintf("user: %s, count: %d", username, count)

	ress := []byte(resval)

	resp.Write(ress)
}
