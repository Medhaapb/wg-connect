package mutex

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
func calcres(index, value int, res []int, mu *sync.Mutex, operation func(int) int, done *int32) {
	result := operation(value)//performs the square operation and store the value in result
	mu.Lock()//locks res slice to allow only on goroutine to write in one time
	res[index] = result//store result in res slice in particular index 
	mu.Unlock()//unlocks res after writing
	atomic.AddInt32(done, 1)//increments done by 1
}
func Mapping(data []int, operation func(int) int) []int {
	res := make([]int, length(data))//create a new slice with same size od input slice to store the result
	var mu sync.Mutex//create an object for mutex
	var done int32// to count number of goroutines that are completed
	for index, value := range data {//iterate through each value
		go calcres(index, value, res, &mu, operation, &done)//call a new goroutine foe each value in the slice
	}
	
	runtime.Goshed()//yield the processor
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
