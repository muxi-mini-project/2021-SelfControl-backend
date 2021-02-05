package main

import "fmt"

func main() {
	var m, n, sum int
	var s [1000]int
	fmt.Scanf("%d%d", &m, &n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &s[i])
	}
	for i := 0; i < m-1; i++ {
		max := i
		for j := i + 1; j < m; j++ {
			if s[j] > s[max] {
				max = j
			}
		}
		s[i], s[max] = s[max], s[i]
	}
	//for i := 0; i < m; i++ {
	//		fmt.Printf("%d ", s[i])

	for i := 0; ; i++ {
		sum += s[i]
		if sum >= n {
			i++
			fmt.Printf("%d\n", i)
			return
		}
	}
}
