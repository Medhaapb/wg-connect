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
