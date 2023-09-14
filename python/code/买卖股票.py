from math import inf
from functools import cache
from typing import List

class Solution:
    # lc 122
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        @cache
        def dfs(i: int, hold: bool) -> int:
            if i < 0:
                return 0 if hold is 0 else -inf
            if hold:
                return max(dfs(i-1, 0)-prices[i], dfs(i-1, 1))
            return max(dfs(i-1, 1)+prices[i], dfs(i-1, 0))
        return dfs(n-1, 0)

    def maxProfitV2(self, prices: List[int]) -> int:
        n = len(prices)
        f = [[0] * 2 for _ in range(n+1)]
        f[0][1] = -inf
        for i, p in enumerate(prices):
            f[i+1][1] = max(f[i][1], f[i][0]-p)
            f[i+1][0] = max(f[i][0], f[i][1]+p)
        
        return f[n][0]