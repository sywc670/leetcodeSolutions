package code

import (
	"fmt"
	"slices"
)

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

// 49. 字母异位词分组
// 本质上是将每个词访问的时间顺序差异给消除，可以通过排序来规定一个顺序，也可以用一个数组甚至一个数来保存，
// 只要这个数可以将字母信息都装下就行
// 1. 判断两个字母异位词通过计数，保存所有26字母列表
// 保存的map中会存有返回值数组的下标来分组
func groupAnagrams(strs []string) (ans [][]string) {
	set := make(map[[26]int]int)
	count := -1
	for _, str := range strs {
		key := [26]int{}
		for _, r := range str {
			key[r-'a']++
		}
		if ret, ok := set[key]; ok { // 如果相同，就放在同一切片中
			ans[ret] = append(ans[ret], str)
			continue
		}
		count++
		set[key] = count
		ans = append(ans, []string{str}) // 如果不同，放在新切片中
	}
	return
}

// 2. 对每个字符串排序
func groupAnagramsV1(strs []string) (ans [][]string) {
	m := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		sortedStr := string(b)
		// map里有没有都无所谓，没有的话append也会自动创建出来
		m[sortedStr] = append(m[sortedStr], str)
	}
	for _, strs := range m {
		ans = append(ans, strs)
	}
	return
}

// 128.最长连续序列
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

func longestConsecutiveV2(nums []int) (ans int) {
	has := make(map[int]bool)
	for _, n := range nums {
		has[n] = true
	}

	for x := range has {
		if has[x-1] {
			continue
		}
		y := x + 1
		for has[y] {
			y++
		}
		ans = max(ans, y-x)
	}
	return
}

// lc 1207. 独一无二的出现次数
func uniqueOccurrences(arr []int) bool {
	set := make(map[int]int)
	newSet := make(map[int]bool)
	for _, n := range arr {
		set[n]++
	}
	for _, v := range set {
		if !newSet[v] {
			newSet[v] = true
			continue
		}
		return false
	}
	return true
}

// opt
func uniqueOccurrencesV2(arr []int) bool {
	set := make(map[int]int)
	newSet := make(map[int]bool)
	for _, n := range arr {
		set[n]++
	}
	for _, v := range set {
		newSet[v] = true
	}
	return len(set) == len(newSet)
}

// lc 1657. 确定两个字符串是否接近
// 种类相同，词频相同，顺序无关
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	set1, set2 := make(map[byte]int), make(map[byte]int)
	for _, w := range word1 {
		set1[byte(w)]++
	}
	for _, w := range word2 {
		set2[byte(w)]++
	}
	for k := range set1 {
		if set2[k] == 0 {
			return false
		}
	}
	// 可以省略，因为词频不同的话，有上面那个判断已经完成了
	// for k := range set2 {
	// 	if set1[k] == 0 {
	// 		return false
	// 	}
	// }
	times1, times2 := make(map[int]int), make(map[int]int)
	for _, v := range set1 {
		times1[v]++
	}
	for _, v := range set2 {
		times2[v]++
	}
	for k, v1 := range times1 {
		if v1 != times2[k] {
			return false
		}
	}
	return true
}

// 2352. 相等行列对
// solve: 只需要rset，不需要cset
// []int无法作为map的key，所以得转化为string类型
// 未掌握
func equalPairs(grid [][]int) (ans int) {
	rset := make(map[string]int)
	for _, r := range grid {
		// wrong: rowStr := string(r)
		rset[fmt.Sprint(r)]++
	}
	for j := range grid {
		col := []int{}
		for i := range grid {
			col = append(col, grid[i][j])
		}
		if count := rset[fmt.Sprint(col)]; count > 0 {
			ans += count
		}
	}
	return
}

// 41. 缺失的第一个正数
// 原地hash，时间复杂度O(n)，空间复杂度O(1)
func firstMissingPositive(nums []int) int {
	// 取值范围为[1,N+1]，故hash规则：map[x-1] = x
	for i := 0; i < len(nums); i++ {
		for 1 <= nums[i] && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	// 置换后连续自然数情况
	return len(nums) + 1
}
