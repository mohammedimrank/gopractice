package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func worker(ctx context.Context, dst chan string, src chan []string) {
	for {
		select {
		case val, ok := <-src:
			if !ok {
				return
			}
			dst <- fmt.Sprintf("out of %v", val)

		case <-ctx.Done():
			return
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var numWorkers int
	var filename string

	flag.StringVar(&filename, "f", "test.csv", "src file")
	flag.IntVar(&numWorkers, "c", 2, "concurrent workers")
	flag.Parse()

	if filename == "" {
		fmt.Println("filename required")
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	src := make(chan []string)
	dest := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Worker %d is executing\n", i)
			worker(ctx, dest, src)
		}(i)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
			}
			src <- record
		}
		close(src)
	}()

	go func() {
		wg.Wait()
		close(dest)
	}()

	for res := range dest {
		fmt.Println(res)
	}

}
