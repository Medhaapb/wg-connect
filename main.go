package main

import (
	"fmt"
	"medhaapb/wg-connect/next"
)

func length(data []int) int {
	c := 0

	for range data {
		c++
	}
	return c
}
func mapper(data []int, operation func(int) int) []int {

	res := make([]int, length(data))
	for i := range data {
		res[i] = operation(data[i])
	}
	return res

}

func square(ele int) int {
	return ele * ele
}

func main() {
	//result := mapper([]int{1, 2, 3, 4, 5, 6}, square)
	//fmt.Println(result)
	//fmt.Println(next.Giveprotocol("https://www.google.com/search?client=ubuntu-sn&hs=Fcn&sca_esv=b597786ab811f977&channel=fs&sxsrf=AHTn8zr3PqJ8aGba7fODIpSs98RTC5KPmg:1741238135135&q=small++coding+exercise+questions+on+%22http%22&sa=X&ved=2ahUKEwiGmo7E2fSLAxUIlq8BHRrHEPEQ5t4CegQIJBAB&biw=1408&bih=703&dpr=1.36"))
	fmt.Println(next.Givedomain("https://www.google.com:0303/search?client=ubuntu-sn&hs=Fcn&sca_esv=b597786ab811f977&channel=fs&sxsrf=AHTn8zr3PqJ8aGba7fODIpSs98RTC5KPmg:1741238135135&q=small++coding+exercise+questions+on+%22http%22&sa=X&ved=2ahUKEwiGmo7E2fSLAxUIlq8BHRrHEPEQ5t4CegQIJBAB&biw=1408&bih=703&dpr=1.36"))

}
