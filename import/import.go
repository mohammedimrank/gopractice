package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var filename = flag.String("f", "REQUIRED", "source CSV file")
var numChannels = flag.Int("c", 4, "num of parallel channels")

//var bufferedChannels = flag.Bool("b", false, "enable buffered channels")

func main() {
	start := time.Now()
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), "\n"))
	if *filename == "REQUIRED" {
		return
	}

	csvfile, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	i := 0
	ch := make(chan []string)
	var wg sync.WaitGroup
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		i++
		wg.Add(1)
		go func(r []string, i int) {
			defer wg.Done()
			fmt.Printf("value of i %d\n", i)
			ch <- r
		}(record, i)

		fmt.Printf("\rgo %d", i)
	}

	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()

	// print channel results (necessary to prevent exit programme before)
	j := 0
	for val := range ch {
		j++
		fmt.Printf("contents of channel %v\n", val)
		// fmt.Printf("\r\t\t\t\t | done %d", j)
	}

	fmt.Printf("\n%2fs", time.Since(start).Seconds())

}
