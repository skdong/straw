package leetcode

import (
	"fmt"
)

func Permute(nums []int){
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var ret [][] int
	for _, num := range nums {
		var new_ret [][] int
		for _, line := range ret{
			for i:=0; i<= len(line); i++{
				new_line := make([]int, len(line)+1)
				copy(new_line, line[:i])
				new_line[i] = num
				copy(new_line[i+1:], line[i:])
				new_ret = append(new_ret, new_line)
			}
		}
		if len(ret) == 0{
			ret = [][]int {{num}}
		}else{
			ret = new_ret
		}
	}
	return ret
    
}