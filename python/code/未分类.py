from typing import List
from functools import reduce

class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        # 334. 递增的三元子序列
        first = second = float('inf')
        for third in nums:
            if third > second:
                return True
            elif third <= first:
                first = third
            else:
                second = third
        return False
    def singleNumber(self, nums: List[int]) -> int:
        # 136. 只出现一次的数字
        return reduce(lambda x, y : x ^ y, nums)
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        # 167. 两数之和 II - 输入有序数组
        left, right = 0, len(numbers)-1
        while left < right:
            sum = numbers[left] + numbers[right]
            if sum == target:
                return [left+1, right+1]
            elif sum > target:
                right = right-1
            else:
                left = left+1
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # 15. 三数之和
        n = len(nums)
        nums.sort()
        res = []
        for i in range(n-2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            if nums[i] + nums[i+1] + nums[i+2] > 0:
                break
            if nums[i] + nums[-2] + nums[-1] < 0:
                continue
            left, right = i+1, n-1
            while left < right:
                sum = nums[i] + nums[left] + nums[right]
                if sum == 0:
                    res.append([nums[i], nums[left], nums[right]])
                    left += 1
                    while left < right and nums[left] == nums[left-1]:
                        left += 1
                    right -= 1
                    while left < right and nums[right] == nums[right+1]:
                        right -= 1
                    continue
                if sum > 0:
                    right -= 1
                else:
                    left += 1
        return res
    def maxArea(self, height: List[int]) -> int:
        # 11. 盛最多水的容器
        left, right = 0, len(height)-1
        maxArea = (right - left) * min(height[left],height[right])
        while left < right:
            if height[left] > height[right]:
                right -= 1
                while left < right and height[right] < height[right+1]:
                    right -= 1
            else:
                left += 1
                while left < right and height[left] < height[left-1]:
                    left += 1
            area = (right - left) * min(height[left], height[right])
            maxArea = max(area, maxArea)
        return maxArea
    def trap(self, height: List[int]) -> int:
        # 42. 接雨水
        ans = left = lmax = rmax = 0
        right = len(height) - 1
        while left < right:
            lmax = max(height[left], lmax)
            rmax = max(height[right], rmax)
            if lmax < rmax:
                ans += lmax - height[left]
                left += 1
            else:
                ans += rmax - height[right]
                right -= 1
        return ans

