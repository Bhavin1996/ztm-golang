package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	index := 0
	i := 1
	if len(nums) == 2 || len(nums) == 1 {
		return len(nums)
	}
	for {
		if nums[i] == nums[i-1] {
			count += 1
			if index == len(nums)-1 || i == len(nums)-1 {
				return int(len(nums))
			}
			i += 1
		} else if nums[i] == nums[i-1] && count == 1 {
			index = i
			count += 1
			if i == len(nums)-1 {
				nums = nums[:i]
				return int(len(nums))
			}
			i += 1
		} else if nums[i] == nums[i-1] && count >= 2 {
			count += 1
			i += 1
		} else if nums[i] != nums[i-1] {
			if count == 0 {
				if index == len(nums)-1 || i == len(nums)-1 {
					return int(len(nums))
				}
				index = 0
				count = 0
				i += 1
			} else if count == 1 {
				if i == len(nums)-1 {
					return int(len(nums))
				}
				index = 0
				count = 0
				i += 1
			} else if count > 1 {
				nums = append(nums[:index], nums[i:]...)
				i = index + 1
				if index == len(nums)-1 || i == len(nums)-1 {
					return int(len(nums))
				}
				index = 0
				count = 0
			}
		}
	}
}
func main() {
	arr := []int{0, 0, 0, 0}
	fmt.Println(removeDuplicates(arr[:]))
}
