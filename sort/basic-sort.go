package sort

import "fmt"

func BubbleSort() {
	a := []int{4, 2, 7, 1, 3}

	for i := len(a) - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	fmt.Println(a)
}

func SelectionSort() {
	a := []int{4, 2, 7, 1, 3}
	var j, i int
	for i = 0; i < len(a); i++ {
		mI := i // minimum index

		for j = i + 1; j < len(a); j++ {
			if a[j] < a[mI] {
				mI = j
			}
		}

		a[i], a[mI] = a[mI], a[i]
	}

	fmt.Println(a)
}

func InsertionSort() {
	a := []int{8, 4, 2, 3}

	var j, i int

	for i = 1; i < len(a); i++ {
		current := a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] > current {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = current
	}

	fmt.Println(a)
}
