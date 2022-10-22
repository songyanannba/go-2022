package main

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)

	for _, v := range nums {
		if v < 0 {
			v = 0 - v
		}
		if nums[l-v] >= 0 {
			nums[l-v] = -nums[l-v]
		}
	}

	re := []int{}
	for i := 1; i <= l; i++ {
		if nums[l-i] > 0 {
			re = append(re, i)
		}
	}
	return re
}
func main()  {
	nums := []int{4,3,2,7,8,2,3,1}
	findDisappearedNumbers(nums)
}
