package main

import "fmt"

func countOfSubsequenceWithSumKHelper(idx, n, count, k, sum int, nums *[]int) int {
	if sum == k {
		count = count + 1
		return count
	}
	if idx == n {
		return 0
	}

	path1 := countOfSubsequenceWithSumKHelper(idx+1, n, count, k, sum+(*nums)[idx], nums)
	path2 := countOfSubsequenceWithSumKHelper(idx+1, n, count, k, sum, nums)

	return path1 + path2
}
func countOfSubsequenceWithSumK(nums []int, k int) int {
	return countOfSubsequenceWithSumKHelper(0, len(nums), 0, k, 0, &nums)
}
func main() {
	nums := []int{4, 9, 2, 5, 1}
	k := 10

	fmt.Println(countOfSubsequenceWithSumK(nums,k))
}