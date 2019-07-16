package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	fmt.Println("请输入有序数组：")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("请输入查找的数：")
	var a int
	fmt.Scanf("\n%d\n", &a)
	i := BinarySearch(arr, n, a)
	if i != -1 {
		fmt.Printf("%d在第%d位置", a, i)
	} else {
		fmt.Printf("%d找不到", a)
	}
}

func BinarySearch(arr []int, n int, a int) int {
	mid := n / 2
	low := 0
	high := n - 1
	for low < high {
		if arr[mid] == a {
			return mid + 1
		} else if arr[mid] > a {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (low + high) / 2
	}
	return -1
}
