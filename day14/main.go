//给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	var nums []int

	numsMap := make(map[int]struct{})
	for _, v := range nums1 {
		numsMap[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := numsMap[v]; ok {
			delete(numsMap, v)
			nums = append(nums, v)
		}
	}

	return nums
}
