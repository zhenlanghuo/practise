package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)
	//fmt.Println(str)
	ans := 0
	for i := 2; i < len(str); i++ {
		temp := str[i] - '0'
		if !(temp <= 9) {
			temp = str[i] - 'A' + 10
		}
		ans = ans*16 + int(temp)
	}
	fmt.Println(ans)
}
