package array

/** leetCode 88
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	c1 := m - 1
	c2 := n - 1
	for c1 >= 0 && c2 >= 0 {
		if nums1[c1] > nums2[c2] {
			nums1[last] = nums1[c1]
			c1--
		} else {
			nums1[last] = nums2[c2]
			c2--
		}
		last--
	}
	//if c2 >= 0 && last >= 0 {
	//	nums1[last] = nums2[c2]
	//	c2--
	//	last--
	//}
	if n >= 0 {
		copy(nums1[0:n+1], nums2[0:n+1])
	}
}
