package channels

import (
	"time"
)
func length(data []int) int {
	c := 0
	for range data {
		c++
	}
	return c
}

func calcres(index int, value int, tempres chan struct{ index, value int }, operation func(int) int) {
	tempres <- struct{ index, value int }{index, operation(value)}
}

func Mapping(data []int, operation func(int) int) []int {
	res := make([]int, length(data))

	resChannel := make(chan struct{ index, value int }, len(data))

	for index, value := range data {
		go calcres(index, value, resChannel, operation)
	}

	for range data {
		result := <-resChannel
		res[result.index] = result.value
	}

	close(resChannel)

	return res
}

func square(ele int) int {
	time.Sleep(1 * time.Second)
	return ele * ele
}
