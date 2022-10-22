package main

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

func main()  {
	climbStairs(5)
}
