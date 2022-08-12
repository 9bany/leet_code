package main

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int

	for _, element := range nums1 {
		m[element]++
	}

	for _, element := range nums2 {
		if m[element] > 0 {
			res = append(res, element)
			m[element]--
		}
	}
	return res
}
