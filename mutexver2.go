package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)
//finds the length of the input slice
func length(data []int) int {
	c := 0
	for range data {
		c++
	}
	return c
}
//calculate the result
func calcres(index int, value int, res []int, operation func(int) int, mu *sync.Mutex, counter *int) {
	mu.Lock()//locks res slice to allow only on goroutine to write in one time
	res[index] = operation(value)//store result in res slice in particular index 
	mu.Unlock()//unlocks res after writing
	*counter++
}
func Mapping(data []int, operation func(int) int) []int {
	res := make([]int, length(data))//create a new slice with same size od input slice to store the result
	var mu sync.Mutex
	var done sync.Mutex
	var counter int
	for index, value := range data {
		go calcres(index, value, res, operation, &mu, &counter, &done)//call a new goroutine foe each value in the slice
	}
	
	// Keep checking if all goroutines are done
	for {
		done.Lock()
		if counter == len(data) {
			done.Unlock()
			break
		}
		done.Unlock()
	}
	return res
}
func square(ele int) int {
	time.Sleep(1 * time.Second)//pauses for 1second
	return ele * ele//returns square of the element
}
func main() {
	result := Mapping([]int{1, 2, 3, 4, 5}, square)
	fmt.Println("Result:", result)
}
