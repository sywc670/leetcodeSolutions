from typing import List

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