package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d%d", &m, &n)
	for i := n - m + 1; i < m+n; i++ {
		if i == m+n-1 {
			fmt.Println(i)
			return
		}
		fmt.Printf("%d ", i)
	}
}
