package dynamic_programming

import(f "fmt")
type Memo map[int]int

func DpTest(){;
	//var a = []int{89,387,89,190,90,10}
	//memo := make(map[int]int);
	//f.Println(fib(6, memo, new(1)))
	//f.Println(dpMax(a, new(1)))
	//f.Println(sumTill100(a, new(1)))
	//f.Println(glob(5,memo));
	quickSortRun()
}

func dpMax(a []int, count *int) int {
	f.Println("Recursion ",*count)
	(*count)++

	if len(a) <= 1 {
		return a[0]
	}

	reminder := dpMax(a[1:], count)

	if a[0] > reminder {
		return a[0]
	}else {
		return reminder
	}
}

func fib(n int,memo Memo, count *int) int {
	f.Println("FIB ",*count)
	(*count)++

	if n == 0 || n == 1 {
		return n
	}

	if _,ok := memo[n]; !ok{
		memo[n] = fib(n-2, memo, count) + fib(n-1, memo, count)
	}

	return memo[n]
}


func sumTill100(a []int, count *int) int {
	f.Println("Sum Till 100 ",*count)
	(*count)++

	if len(a) <= 0{
		return 0
	}

	remaining := sumTill100(a[1:], count)

	if (a[0] + remaining) > 100{
		return remaining
	}else {
		return a[0] + remaining
	}
}

func glob(n int, memo Memo) int {
	if n == 1 {
		return 1
	}

	if _,ok := memo[n]; !ok{
		memo[n] = 1 + glob(n- glob(glob(n-1, memo), memo), memo)
	}

	return memo[n]
}

// Quick sort

func quickSortRun() {
	a := &[5]int{4,6,1,3,2}
	f.Println(a)
	quickSort(a, 0, len(a)-1)
	f.Println(a)
}

func quickSort(a *[5]int, lo, hi int)  {
	if lo >= hi{
		return
	}

	pivot := partition(a, lo, hi)

	quickSort(a, lo, pivot-1)
	quickSort(a, pivot+1, hi)
}

func partition(a *[5]int, lo int, hi int) int {
	pivotIndex := hi
	pivot := a[pivotIndex]

	hi -= 1
	for {

		for a[lo] < pivot{
			lo++
		}

		for a[hi] > pivot{
			hi--
		}

		if lo >= hi {
			break
		}else {
			a[lo], a[hi] = a[hi], a[lo]
		}

	}

	a[lo], a[pivotIndex] = a[pivotIndex], a[lo]

	return lo

}