package mutex

import (
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

func calcres(index, value int, res []int, mu *sync.Mutex, operation func(int) int) {

	result := operation(value)

	mu.Lock()
	res[index] = result
	mu.Unlock()
}

func Mapping(data []int, operation func(int) int) []int {
	res := make([]int, length(data))
	var mu sync.Mutex

	for index, value := range data {
		go calcres(index, value, res, &mu, operation)
	}
	time.Sleep(2 * time.Second)

	return res
}

func square(ele int) int {
	time.Sleep(1 * time.Second)

	return ele * ele
}
