package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	result := 0
	pre := 2
	prePre := 1
	for i := 3; i <= n; i++ {
		result = pre + prePre
		prePre = pre
		pre = result
	}

	return result
}

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs1(n-1) + climbStairs1(n-2)


}

func main()  {
	stairs := climbStairs(5)
	fmt.Println(stairs)


	stairs1 := climbStairs1(5)
	fmt.Println(stairs1)
}
