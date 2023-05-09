package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

func getComic(comicID int) (comic *Comic, err error) {
	url := fmt.Sprintf(baseXkcdURL, comicID)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// body, err := ioutil.ReadAll(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return comic, nil
}

func main() {

	comicsNeeded := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	comicMap := make(map[int]*Comic, len(comicsNeeded))
	var wg sync.WaitGroup
	for _, id := range comicsNeeded {
		fmt.Println(id)
		wg.Add(1)
		go func(id int) {
			comic, err := getComic(id)

			if err != nil {
				return
			}
			comicMap[id] = comic

			wg.Done()
		}(id)
	}
	fmt.Printf("comic details %v ", comicMap)
	// wg.Wait()
}
