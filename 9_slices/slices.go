package main

import (
	"fmt"
)

func main(){
	// null == nil (in go)
    // uninitialized slice in go is nil
    // var nums[] int
    // fmt.Println(nums)
    // fmt.Println(nums == nil)
    // fmt.Println(len(nums))

    // to not inilize with nil
    // var nums = make([]int, 0, 5)     // type, size, capacity
    //or
    // nums := []int{}

    // fmt.Println(arr)
    // capacity -> max number of ele that can be stored for now
    // fmt.Println(cap(arr))
    // nums = append(nums,1)
    // nums = append(nums,2)
    // nums = append(nums,3)

    // copy function
    // var nums = make([]int, 0, 5)
    // nums = append(nums, 1)
    
    // var nums2 = make([]int, len(nums))
    
    // copy(nums2,nums)
    // fmt.Println(nums, nums2)

    //slice operator
    // var nums = [] int {1,2,3,4,5,6,7}
    // fmt.Println(nums[:3])
    // fmt.Println(nums[0:3])
    // fmt.Println(nums[3:5])
    // fmt.Println(nums[3:])

    // slice
    // var nums1 = []int {1,2}
    // var nums2 = []int {1,2}

    // fmt.Println(slices.Equal(nums1,nums2))

    // 2d slice
    var nums = [][] int {{1,2,3},{4,5,6}}
    fmt.Println(nums)
}