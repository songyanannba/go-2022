package main

import "fmt"

func solve(s string, t string) string {
	m := len(s)
	n := len(t)
	if m > n {
		m, n = n, m
		s, t = t, s
	}
	ret := make([]byte, n+1)
	carry := 0
	for i := 0; i < n; i++ {
		longN := int(t[n-i-1] - '0')
		shortN := 0
		if m-i-1 >= 0 {
			shortN = int(s[m-i-1] - '0')
		}
		result := longN + shortN + carry
		carry = result / 10
		result = result % 10
		ret[n-i] = byte(result + '0')
	}
	fmt.Println(ret)

	if carry == 1 {
		ret[0] = '1'
		return string(ret)
	}
	return string(ret[1:])
}

func main() {
	str := "1003"
	str1 := "222"
	s := solve(str, str1)
	fmt.Println(s)
}