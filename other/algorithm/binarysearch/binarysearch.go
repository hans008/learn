package main

import "fmt"

func BinarySearch(arr []int,target int) int {
	//定义边界，[l,r] 左闭右闭的区间
	n := len(arr)

	l,r :=0,n-1

	for l <= r {
		mid := (l+r)/2

		if target == arr[mid] {
			return mid
		}

		if target > arr[mid] {
			l = mid + 1
		} else if target < arr[mid] {
			r = mid - 1
		}
	}
	return -1
}

func main()  {
	arr :=[]int{1,2,5,8,10,15,30,32}
	target := 8

	result := BinarySearch(arr,target)
	fmt.Println(result)
}


