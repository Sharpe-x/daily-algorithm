package main

/*
https://leetcode.cn/problems/remove-element/

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

// O(n^2)
func removeElementByOn(nums []int, val int) int {

	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			for j := i + 1; j < length; j++ {
				nums[j-1] = nums[j]
			}
			i--
			length--
		}
	}
	return length
}

// 快慢指针
// 快指针用来寻找新数组的元素
// 慢指针用来记录新数组的最后一个位置
// 最后慢指针的长度就是新数组的大小
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow = slow + 1
		}
	}
	return slow
}
