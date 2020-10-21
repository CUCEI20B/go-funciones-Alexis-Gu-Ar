package main

import "fmt"

func swap(s []int64, i int, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func Burbuja(s []int64) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j, j+1)
			}
		}
	}
}

func main() {
	var s = []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 11}

	Burbuja(s)
	fmt.Print(s)
}
