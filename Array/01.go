package array


// 1. Find the missing number in a given array of size n containing numbers from 1 to n+1.
func MissingNumber(nums []int) int {
	// sum of all
	expectedSum := (len(nums) +1) * (len(nums)+2)/ 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

// 2. Reverse array inplace
func ReverseArray(nums []int ) []int {
	pt1, pt2 := 0, len(nums) -1
	for pt1 < pt2 {
		nums[pt1], nums[pt2] = nums[pt2], nums[pt1]
		pt1++
		pt2--
	}
	return nums
}
