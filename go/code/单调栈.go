package code

import "sort"

// 思路：单调栈题目，首先有一串数字，当数字是单调减少时，入栈，
// 当遇到比栈顶更大的数字时，出栈之前的一些数字，并做一些工作
// 遍历顺序和单调顺序可能变化

// 739. 每日温度
// 单调栈每次都清除无用数据，保持元素有序
// 从右到左
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= t {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

// 从左到右
func dailyTemperaturesV1(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// lc 42
// 思路：横向计算水面积，需要左边最近的顶和右边最近的顶以及中间的底
// 从左到右遍历，遇到比栈顶大，出栈并计算
func trapV1(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= h { // 当高度连续时也成立，因为后续相同高度面积为0
			bottomHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			width := i - stack[len(stack)-1] - 1
			ans += width * (min(height[stack[len(stack)-1]], h) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return
}

// lc 面试题 17.08： 马戏团人塔
// 这个题可以看作二维LIS，排序，使得题目变成一维LIS
func bestSeqAtIndex(height []int, weight []int) int {
	l := len(height)
	type person struct {
		height, weight int
	}
	// 也可以用二维数组来做
	personSlice := make([]person, l)
	for i, h := range height {
		personSlice[i] = person{height: h, weight: weight[i]}
	}
	// 身高升序排序，体重降序排序，其中体重降序是为了解决身高重复情况
	// 体重升序会让身高重复的人被错误计算在内
	sort.Slice(personSlice, func(i, j int) bool {
		if personSlice[i].height == personSlice[j].height {
			return personSlice[i].weight > personSlice[j].weight
		}
		return personSlice[i].height < personSlice[j].height
	})

	// LIS
	// 单调栈+二分查找
	mStack := make([]int, 0, l/2)
	mStack = append(mStack, personSlice[0].weight) // 第一个提前加入
	for i := 1; i < l; i++ {
		target := personSlice[i].weight
		if target > mStack[len(mStack)-1] {
			mStack = append(mStack, target)
		} else if target < mStack[len(mStack)-1] {
			left, right := -1, len(mStack)
			for left+1 < right {
				mid := (left + right) / 2
				if target > mStack[mid] { // SOLVE: 二分比较成mid下标了
					left = mid
				} else {
					right = mid
				}
			}
			// 替代比target大的第一个数
			// 因为比target大的其它数还会影响结果，所以不能删除
			// 如[1 2 0 3] 0加入不能删除1和2 应为[0 2 3]
			// 只是替换还可以使得所求长度就是单调栈高度
			mStack[right] = target
		}
	}
	return len(mStack)
}

// 单调栈+二分查找
func lengthOfLISV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mStack := make([]int, 0)
	mStack = append(mStack, nums[0])
	for i := 1; i < len(nums); i++ {
		target := nums[i]
		if target > mStack[len(mStack)-1] {
			mStack = append(mStack, target)
		} else if target < mStack[len(mStack)-1] {
			left, right := -1, len(mStack)
			for left+1 < right {
				mid := left + (right-left)>>1
				if target <= mStack[mid] { // 等于也需要替换
					right = mid // right是比target大的第一个
				} else {
					left = mid
				}
			}
			mStack[right] = target
		}
	}
	return len(mStack)
}

// 901. 股票价格跨度
type StockSpanner struct {
}

func ConstructorStock() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	return 0
}
