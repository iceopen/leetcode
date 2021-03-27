package _83_move_zeroes

func moveZeroes(nums []int) {
	var num = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			num++
		} else if num != 0 {
			nums[i-num] = nums[i]
			nums[i] = 0
		}
	}
}
