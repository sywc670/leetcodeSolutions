# lc 1143
class Solution:
    def longestCommonSubsequence(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        f = [[0] * (m + 1) for _ in range(n + 1)]
        for i, x in enumerate(s):
            for j, y in enumerate(t):
                f[i + 1][j + 1] = f[i][j] + 1 if x == y else \
                                  max(f[i][j + 1], f[i + 1][j])
        return f[n][m]

print(Solution().longestCommonSubsequence("abcde", "ace"))