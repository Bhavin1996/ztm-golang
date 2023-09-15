package main

import "fmt"

func removeDuplicates(nums []int) (x int, arr []int) {
	count := 0
	index := 0
	i := 1
	if len(nums) == 2 || len(nums) == 1 {
		return len(nums), nums[:]
	}
	for {
		if nums[i] == nums[i-1] {
			count += 1
			if count == 2 {
				index = i
				if i == len(nums)-1 {
					nums = nums[:i]
					return len(nums), nums[:]
				}
				i += 1
			} else if count > 2 {
				if i == len(nums)-1 {
					nums = nums[:index]
					return len(nums), nums[:]
				}
				i += 1
			} else {
				if i == len(nums)-1 {
					nums = nums[:i+1]
					return len(nums), nums[:]
				}
				i += 1
			}
		} else if nums[i] != nums[i-1] {
			if count == 0 {
				if i == len(nums)-1 {
					return len(nums), nums[:]
				}
				i += 1
			} else if count == 1 {
				if i == len(nums)-1 {
					nums = nums[:i+1]
					return len(nums), nums[:]
				}
				nums = append(nums[:i], nums[i:]...)
				i = i + 1
				count = 0
			} else {
				nums = append(nums[:index], nums[i:]...)
				if i == len(nums)-1 {
					return len(nums), nums[:]
				}
				count = 0
				i = index + 1
			}
		}
	}
}

func main() {
	arr := []int{1, 1, 1, 2}
	fmt.Println(removeDuplicates(arr[:]))
}
