package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func length(data []int) int {
	c := 0
	for range data {
		c++
	}
	return c
}

func calcres(index, value int, res []int, mu *sync.Mutex, operation func(int) int, done *int32) {
	result := operation(value)

	mu.Lock()
	res[index] = result
	mu.Unlock()
	atomic.AddInt32(done, 1)
}

func Mapping(data []int, operation func(int) int) []int {
	res := make([]int, length(data))
	var mu sync.Mutex
	var done int32

	for index, value := range data {
		go calcres(index, value, res, &mu, operation, &done)
	}

	startTime := time.Now()
	timeoutDuration := 5 * time.Second 
	fmt.Println(startTime, timeoutDuration)
	for {

		if atomic.LoadInt32(&done) == int32(length(data)) {
			return res
		}

		if time.Since(startTime) > timeoutDuration {
			break
		}

	}
	return nil
}

func square(ele int) int {
	time.Sleep(1 * time.Second)
	return ele * ele
}
