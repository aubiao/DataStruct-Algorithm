package main

import "fmt"

func Swap(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

func PreDown(arr []int, i int, N int) {
	var child int
	var temp int
	for temp = arr[i]; 2*i+1 < N; i = child {
		child = 2*i + 1
		if child != N-1 && arr[child] < arr[child+1] {
			child++
		}
		if temp < arr[child] {
			arr[i] = arr[child]
		} else {
			break
		}
	}
	arr[i] = temp
}

func HeapSort(arr []int, N int) {
	var i int
	for i = N / 2; i >= 0; i-- {
		PreDown(arr, i, N)
	}
	for i = N - 1; i > 0; i-- {
		Swap(&arr[0], &arr[i])
		PreDown(arr, 0, i)
	}
}

func main() {
	var num int
	fmt.Println("请输入数组长度:")
	fmt.Scanf("%d\n", &num)
	arr := make([]int, num) //Go的数组不能用变量去定义长度，只能用常量，所以这里用切片
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	HeapSort(arr, num)
	fmt.Println(arr)
	fmt.Println(cap(arr))
}
