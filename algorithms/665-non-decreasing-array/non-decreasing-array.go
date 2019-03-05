package nonDecreasingArray

import "sort"

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pre := deepCopy(nums)
			pre[i-1] = pre[i]
			next := deepCopy(nums)
			next[i] = next[i-1]
			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}
	return true
}

func deepCopy(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

// func checkPossibility(nums []int) bool {
// 	var l = len(nums)-1
// 	var n = 0
// 	for i:=0;i<l;i++ {
// 		if nums[i] > nums[i+1] {
// 			n ++
// 			if i>0 && nums[i-1] > nums[i+1]{
// 				nums[i+1] = nums[i]
// 			}
// 		}
// 	}
// 	return n<=1
// }
