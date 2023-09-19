package code

import "math"

// 643. 子数组最大平均数 I
// 超时：使用滑窗
func findMaxAverage(nums []int, k int) (ans float64) {
	ans = -math.MaxFloat64
	for left := range nums {
		if left+k > len(nums) {
			return
		}
		var sum int
		for i := 0; i < k; i++ {
			sum += nums[left+i]
		}
		res := float64(sum) / float64(k)
		ans = max(ans, res)
	}
	return
}

func findMaxAverageV2(nums []int, k int) (ans float64) {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for right := range nums {
		if right < k {
			continue
		}
		sum = sum - nums[right-k] + nums[right]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

// 1493. 删掉一个元素以后全为 1 的最长子数组
func longestSubarray(nums []int) (ans int) {
	left, ones := 0, 0
	var zeroFlag bool

	for right, v := range nums {
		if v == 1 {
			// last one still no zero
			if right == len(nums)-1 && !zeroFlag {
				return right
			}
			ones++
			ans = max(ans, ones)
			continue
		}

		// when v == 0
		if !zeroFlag {
			zeroFlag = true
		} else {
			// move left
			for {
				if nums[left] == 1 {
					ones--
					left++
				} else {
					left++
					break
				}
			}
		}
	}
	return
}
