package main

import "fmt"

func combinationSumHelper(
	idx, n, sum, target int,
	nums, list []int,
	ans *[][]int,
) {
	if target == sum {
		temp := make([]int, len(list))
		copy(temp, list)
		*ans = append(*ans, temp)
		return
	}

	if idx == n || sum > target {
		return
	}

	list = append(list, nums[idx])
	combinationSumHelper(idx, n, sum+nums[idx], target, nums, list, ans)
	list = list[:len(list)-1]
	combinationSumHelper(idx+1, n, sum, target, nums, list, ans)

}

func combinationSum(target int, nums []int) [][]int {
	var ans [][]int
	var list []int
	combinationSumHelper(0, len(nums), 0, target, nums, list, &ans)
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5}
	target := 7
	ans := combinationSum(target,nums)
	for i,ans := range(ans){
		fmt.Println(i+1,ans)
	}
}