package code

// lc 136. 只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// 338. 比特位计数
// dp
// 奇数：二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1。
// 偶数：二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多。因为最低位是 0，除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的。
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i%2 == 1 {
			bits[i] = bits[i-1] + 1
		} else {
			bits[i] = bits[i/2]
		}
	}
	return bits
}

// i&1如果是奇数就为1，偶数就为0
func countBitsV2(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}
