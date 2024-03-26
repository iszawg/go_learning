package main

import "fmt"

func merge_sort (q []int, l, r int) {

}

func main(){
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	merge_sort(q, 0, n - 1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
	return
}