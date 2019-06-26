package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := 0; i < num-1; i++ {
		flag := false
		for j := num - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	fmt.Println(arr)
}
