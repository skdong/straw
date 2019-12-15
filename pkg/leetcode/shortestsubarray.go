package leetcode

func shortestSubArray(A []int, K int) int {
	var ret int = -1
	var max, maxIndex, sum, count int

	for i := range A {
		for j := maxIndex; j < len(A); j++ {
			sum += A[j]
			count++
			if sum > max{
				max = sum
				maxIndex = i+count
			}
			if sum >= K && count > ret {
				ret = count
			}
		}
		sum = max - A[i]
		count = maxIndex - i
	}
	return ret
}
