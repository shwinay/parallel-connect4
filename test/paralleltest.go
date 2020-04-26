package main

import "fmt"

type Result struct {
	first  int
	second int
}

func sendOverChan(i int, c chan Result) {
	c <- Result{i, i}
}

func main() {

	c := make(chan Result)
	for i := 0; i < 10; i++ {
		go sendOverChan(i, c)
	}

	results := make([]Result, 0)

	for i := 0; i < 10; i++ {
		r := <-c
		results = append(results, r)
	}

	for _, res := range results {
		fmt.Println(res)
	}

}
