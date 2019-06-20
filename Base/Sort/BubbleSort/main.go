package BubbleSort

import "fmt"

func main() {
	ar := [10]int{9, 8, 6, 4, 2, 7, 1, 3, 0, 5}
	num := len(ar)
	for i := 0; i < num; i++ {
		flag := false
		for j := i + 1; j < num; j++ {
			if ar[i] < ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	fmt.Println(ar)
}
