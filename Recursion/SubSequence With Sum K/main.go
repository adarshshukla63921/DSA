package main

import "fmt"

func subsequenceWithSumKHelper(idx, n, k, sum int, nums *[]int) bool {
	if sum == k {
		return true
	}

	if sum < 0 || idx == n {
		return false
	}

	path1 := subsequenceWithSumKHelper(idx+1, n, k, sum+(*nums)[idx], nums)
	path2 := subsequenceWithSumKHelper(idx+1, n, k, sum, nums)

	return path1 || path2

}

func subsequenceWithSumK(nums []int, k int) bool {
	return subsequenceWithSumKHelper(0, len(nums), k, 0, &nums)
}

func main() {
	nums := []int{1, 4, 2,8}

	k := 16

	fmt.Println(subsequenceWithSumK(nums,k))
}