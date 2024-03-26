package main

import "fmt"

func merge_sort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	merge_sort(q, l, mid)
	merge_sort(q, mid+1, r)
	res := make([]int, r-l+1)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			res[k] = q[i]
			k++
			i++
		} else {
			res[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		res[k] = q[i]
		k++
		i++
	}
	for j <= r {
		res[k] = q[j]
		k++
		j++
	}
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = res[j]
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	merge_sort(q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
	return
}
