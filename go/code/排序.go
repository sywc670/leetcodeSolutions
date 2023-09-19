package code

// 快速排序
func findKthLargestV1(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// 堆排序
func findKthLargestV2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeapV2(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapifyV2(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeapV2(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapifyV2(a, i, heapSize)
	}
}

func maxHeapifyV2(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapifyV2(a, largest, heapSize)
	}
}

// 冒泡排序
func BubbleSort(nums []int) {
	// 控制外圈循环范围逐渐缩小
	for i := 0; i < len(nums)-1; i++ {
		// j代表开始交换的下标
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// 选择排序
func SelectSort(nums []int) {
	// 外圈表示循环的范围
	for i := 0; i < len(nums); i++ {
		// key是每次循环的最小值下标
		key := i
		// 内圈是遍历
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[key] {
				key = j
			}
		}
		if key != i {
			nums[i], nums[key] = nums[key], nums[i]
		}
	}
}

// 插入排序
func InsertSort(nums []int) {
	// 外层循环枚举有序列表后一位下标
	for i := 1; i < len(nums); i++ {
		// 内层来控制交换过程
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

// 希尔排序
func ShellSort(nums []int) {
	n := len(nums)
	gap := 1
	for gap < n/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		// 外层枚举有序子序列后一位
		for i := gap; i < n; i++ {
			// 内层枚举前一位
			for j := i; j-gap >= 0 && nums[j-gap] > nums[j]; j -= gap {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
			}
		}
		gap /= 3
	}
}

// 归并排序
func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	key := n / 2
	left := MergeSort(nums[:key])
	right := MergeSort(nums[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}

// 快速排序
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		key := (left + right) / 2
		i, j := left, right
		for i <= j {
			for nums[i] < nums[key] {
				i++
			}
			for nums[j] > nums[key] {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}

		if left < j {
			quickSort(nums, left, j)
		}
		if right > i {
			quickSort(nums, i, right)
		}
	}
}

// 堆排序
func HeapSort(nums []int) {
	n := len(nums)
	buildMaxHeap(nums, n)
	heapSort(nums)
}

func heapSort(nums []int) {
	for last := len(nums) - 1; last > 0; {
		nums[0], nums[last] = nums[last], nums[0]
		last--
		// 下滤
		downMaxHeapify(nums, 0, last+1)
	}
}
