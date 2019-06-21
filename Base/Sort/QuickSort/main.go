package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	arr := make([]int, num, 1000)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	QuickSort(arr, 0, num-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		a := Partition(arr, low, high)
		QuickSort(arr, low, a-1)
		QuickSort(arr, a+1, high)
	}
}

func Partition(arr []int, low int, high int) int {
	p := arr[low]
	for low < high {
		for low < high && arr[high] >= p {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= p {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = p
	return low
}
