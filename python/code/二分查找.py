from bisect import bisect_left
from typing import List

# 2300. 咒语和药水的成功对数
class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        return [len(potions) - bisect_left(potions, success / s) for s in spells]