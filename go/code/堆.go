package code

import (
	"container/heap"
	"slices"
	"sort"
)

// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		downMaxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		downMaxHeapify(a, i, heapSize)
	}
}

// 215中这个下滤其实不用考虑i出界，因为builidMaxHeap中i取值保证了不会出界
// 这里是为了更通用加上的
func downMaxHeapify(nums []int, i, heapSize int) {
	if i >= heapSize {
		return
	}
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		downMaxHeapify(nums, largest, heapSize)
	}
}

func downMinHeapify(nums []int, i, heapSize int) {
	if i >= heapSize {
		return
	}
	l, r, smallest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] < nums[smallest] {
		smallest = l
	}
	if r < heapSize && nums[r] < nums[smallest] {
		smallest = r
	}
	if smallest != i {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		downMinHeapify(nums, smallest, heapSize)
	}
}

// 2336. 无限集中的最小数字
type SmallestInfiniteSet struct {
	min  int
	heap []int
	// heapHash用于AddBack时快速找到堆中是否有这个值
	heapHash map[int]struct{}
}

func ConstructorV() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		min:      1,
		heap:     []int{},
		heapHash: map[int]struct{}{},
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	m := s.min
	// 搜索堆中有没有比m更小的值
	if len(s.heap) > 0 && s.heap[0] < m {
		m = s.heap[0]
		//下滤
		s.heap[0] = s.heap[len(s.heap)-1]
		s.heap = s.heap[:len(s.heap)-1]

		downMinHeapify(s.heap, 0, len(s.heap))
		delete(s.heapHash, m)
		return m
	}
	s.min++
	return m
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num >= s.min {
		return
	}
	_, ok := s.heapHash[num]
	if ok {
		return
	} else {
		s.heapHash[num] = struct{}{}
	}
	s.heap = append(s.heap, num)
	upMinHeapify(s.heap, len(s.heap)-1)
}

// 不需要heapSize来判断边界，因为是上滤
func upMinHeapify(nums []int, i int) {
	parent := (i - 1) / 2
	if parent == i {
		return
	}
	if nums[parent] > nums[i] {
		nums[parent], nums[i] = nums[i], nums[parent]
		upMinHeapify(nums, parent)
	}
}

// 不需要heapSize来判断边界，因为是上滤
func upMaxHeapify(nums []int, i int) {
	parent := (i - 1) / 2
	if parent == i {
		return
	}
	if nums[parent] < nums[i] {
		nums[parent], nums[i] = nums[i], nums[parent]
		upMinHeapify(nums, parent)
	}
}

// 2542. 最大子序列的分数
// heap可以自己实现
func maxScore(nums1, nums2 []int, k int) int64 {
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}
	// 对下标排序，不影响原数组的顺序
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })

	h := hp{make([]int, k)}
	sum := 0
	for i, idx := range ids[:k] {
		sum += nums1[idx]
		h.IntSlice[i] = nums1[idx]
	}
	heap.Init(&h)

	ans := sum * nums2[ids[k-1]]
	for _, i := range ids[k:] {
		x := nums1[i]
		if x > h.IntSlice[0] {
			sum += x - h.replace(x)
			ans = max(ans, sum*nums2[i])
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Push(any)            {}
func (hp) Pop() (_ any)        { return }
func (h hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }

func totalCost(costs []int, k int, candidates int) (ans int64) {
	n := len(costs)
	if candidates*2+k > n {
		slices.Sort(costs)
		for _, v := range costs[:k] {
			ans += int64(v)
		}
		return
	}
	pre := hp{costs[:candidates]}
	suf := hp{costs[n-candidates:]}
	heap.Init(&pre)
	heap.Init(&suf)

	for i, j := candidates, n-1-candidates; k > 0; k-- {
		if pre.IntSlice[0] <= suf.IntSlice[0] {
			ans += int64(pre.replace(costs[i]))
			i++
		} else {
			ans += int64(suf.replace(costs[j]))
			j--
		}
	}
	return
}
