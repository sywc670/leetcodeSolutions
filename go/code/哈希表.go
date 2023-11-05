package code

// lc 1
// 与167题区别在于无序，且要返回下标所以不能排序
// 思路在于两个数相加变成nums[i] == target - nums[j]
// 遍历每个数可以存下之前数的值和下标，之后每个数再查找需要的数是否存在
func twoSumV(nums []int, target int) []int {
	index := make(map[int]int)
	for i, x := range nums {
		if j, ok := index[target-x]; ok {
			return []int{j, i}
		}
		index[x] = i
	}
	return nil
}

// lc 49 字母异位词分组
// 1. 判断两个字母异位词通过计数，保存所有26字母列表
// 2. 对每个字符串排序
// 保存的map中会存有返回值数组的下标来分组
func groupAnagrams(strs []string) (ans [][]string) {
	set := make(map[[26]int]int)
	count := -1
	for _, str := range strs {
		key := [26]int{}
		for _, r := range str {
			key[r-'a']++
		}
		if ret, ok := set[key]; ok {
			ans[ret] = append(ans[ret], str)
			continue
		}
		count++
		set[key] = count
		ans = append(ans, []string{str})
	}
	return
}

// lc 128 最长连续序列
// 存哈希表，然后遍历左右算长度，超时了
func longestConsecutive(nums []int) (ans int) {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		left, right := n, n
		for {
			_, ok := set[right+1]
			if !ok {
				break
			}
			right++
		}
		for {
			_, ok := set[left-1]
			if !ok {
				break
			}
			left--
		}
		ans = max(ans, right-left+1)
		set[n] = struct{}{}
	}
	return
}

// 哈希集合
// 先遍历一遍存表，第二遍如果有小于该数的数就不进行判断，从而优化时间，但还是超时
func longestConsecutiveV1(nums []int) (ans int) {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for _, n := range nums {
		if _, ok := set[n-1]; ok {
			continue
		}
		right := n
		for {
			_, ok := set[right+1]
			if !ok {
				break
			}
			right++
		}
		ans = max(ans, right-n+1)
	}
	return
}
