package next

import (
	"strings"
)

func Giveprotocol(url string) string {
	index := strings.Index(url, ":")
	if index == -1 {
		return url
	}
	return url[:index]

}

// https://www.google.com/search?client=ubuntu-sn&hs=Fcn&sca_esv=b597786ab811f977&channel=fs&sxsrf=AHTn8zr3PqJ8aGba7fODIpSs98RTC5KPmg:1741238135135&q=small++coding+exercise+questions+on+%22http%22&sa=X&ved=2ahUKEwiGmo7E2fSLAxUIlq8BHRrHEPEQ5t4CegQIJBAB&biw=1408&bih=703&dpr=1.36
func Givedomain(url string) string {
	index1 := strings.Index(url, "://")
	temp := url[index1+3:]
	index2 := strings.Index(temp, "/")
	temp = temp[:index2]
	index3 := strings.Index(temp, ":")
	if index3 != -1 {
		temp = temp[:index3]
	}
	index4 := strings.Index(temp, ".")
	temp = temp[index4+1:]
	return temp

}
