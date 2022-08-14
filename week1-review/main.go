package main

func main() {

}

// 左闭右闭区间
func search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {

		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
