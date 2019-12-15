package sort

func sort(nums []int, i int, j int) {
	// sort
	k := i
	tmp := nums[k]
	left := i
	right := j-1
	for left != right{
		if nums[right] < tmp{
			nums[k] = nums[right]
			k = right
			left ++
		}else if nums[left] > tmp{
			nums[k] = nums[left]
			k = left
			right --
		}else if k==left{
			right --
		}else if k==right{
			left ++
		}
	}
	nums[k] = tmp
	if k-i > 1 {
		sort(nums, i, k)
	}
	if j-k-1 > 1 {
		sort(nums, k+1, j)
	}

}

func QuickSort(nums []int) []int {
	sort(nums, 0, len(nums))
	return nums
}
