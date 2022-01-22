package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func getLen(s string, left, right int) int {
	size := len(s)
	for left >= 0 && right < size && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func getLongestPalindrome(A string) int {
	size := len(A)
	var ans int

	for i := 0; i < size; i++ {
		left := getLen(A, i-1, i+1)
		right := getLen(A, i, i+1)
		ans  = max(ans ,max(left ,right))
	}
	return ans
}



func getLongestPalindrome1(A string) int {
	var b = []byte(A)
	var max int
	fmt.Println(b)
	for i := 0 ; i< len(b) ; i++ {
		var left int = i
		var right int = i
		var count int

		if b[left] == b[right] {
			count = -1
			for left >= 0 && right < len(b) {
				if b[left] ==b[right] {
					count+= 2
					if count >max {
						max = count
					}
					left--
					right++
				} else {
					break
				}
			}
		}
		count = 0
		left= i
		right = i
		if right + 1 < len(b) && b[left] == b[right+1] {
			right = right + 1
			for left >= 0 && right < len(b){
				if b[left] == b[right] {
					count += 2
					if count > max {
						max = count
					}
					left--
					right++
				} else {
					break
				}
			}
		}
	}
	return  max
}



func main() {
	var a = "ababc"
	fmt.Println(getLongestPalindrome(a))
	fmt.Println(getLongestPalindrome1(a))
}