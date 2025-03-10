package main

import (
	"fmt"
	"medhaapb/wg-connect/channels"
	"sync"
	"time"
)

func length(data []int) int {
	c := 0

	for range data {
		c++
	}
	return c
}
func calcres(index int, value int, res []int, operation func(int) int, wg *sync.WaitGroup) {
	fmt.Println(&wg)
	defer wg.Done()
	res[index] = operation(value)
}

func mapper(data []int, operation func(int) int) []int {

	res := make([]int, length(data))
	var wg sync.WaitGroup
	for index, value := range data {
		wg.Add(1)
		fmt.Println(&wg)
		go calcres(index, value, res, operation, &wg)

	}
	wg.Wait()
	return res

}

func square(ele int) int {
	time.Sleep(1 * time.Second)

	return ele * ele
}

func main() {
	tm1 := time.Now()
	result := mapper([]int{1, 2, 3, 4, 5, 6}, square)
	fmt.Println(result)
	tm2 := time.Now()
	fmt.Println(tm1, tm2)
	fmt.Println(channels.Mapping([]int{1, 2, 3, 4, 5}, square))

}
