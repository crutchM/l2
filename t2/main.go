package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "asfd3/42/52"
	fmt.Println(unpack(str))
}

//ничего лучше чем добавить флаги для "экранирования" чисел я не придумал, говнокод но работает
func unpack(line string) string {
	symbols := []rune(line)
	result := ""
	flag := false
	f2 := false
	if len(symbols) == 0 || unicode.IsDigit(symbols[0]) {
		return "incorrect row"
	}
	for i, v := range symbols {
		if string(v) == "/" {
			flag = true
			continue
		}
		if flag {
			if i == len(symbols)-1 {
				result += string(v)
				break
			}
			if unicode.IsDigit(symbols[i+1]) {
				n, _ := strconv.Atoi(string(symbols[i+1]))
				for j := 0; j < n; j++ {
					result += string(v)
				}
				flag = false
				f2 = true

			} else {
				result += string(v)
				flag = false
				continue
			}
		} else {
			if f2 {
				f2 = false
			} else {
				if unicode.IsDigit(v) {
					if unicode.IsDigit(symbols[i-1]) {
						return "incorrect row"
					}
					n, _ := strconv.Atoi(string(v))
					for j := 0; j < n; j++ {
						result += string(symbols[i-1])
					}
				} else {
					result += string(v)
				}
			}
		}
	}

	return result
}
