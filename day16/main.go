package main

//给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

func main() {

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	suma := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			suma[nums1[i]+nums2[j]]++
		}
	}

	ans := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			need := 0 - (nums3[i] + nums4[j])
			ans += suma[need]
		}
	}

	return ans
}
