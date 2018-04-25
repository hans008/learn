package main

import "fmt"

/*
https://leetcode.com/problems/move-zeroes/description/
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

func moveZero(arr []int) []int{
	var nonZeroArry []int

	for i :=0;i<len(arr);i++ {
		if arr[i] != 0 {
			nonZeroArry = append(nonZeroArry,arr[i])
		}
	}

	for i :=0;i<len(nonZeroArry);i++ {
		arr[i] = nonZeroArry[i]
	}

	for i :=len(nonZeroArry);i<len(arr); i++{
		arr[i] = 0
	}

	return arr
}


func main(){
	arr :=[]int{0,1,0,3,12}

	a :=moveZero(arr)

	for i :=0;i<len(a);i++{
		fmt.Println(a[i])
	}
}

