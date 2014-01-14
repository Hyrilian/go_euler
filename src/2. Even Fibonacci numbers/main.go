package main

import "fmt"

func fibonacci() int64 {
	
	var fib_array = make([]int64, 40)
	var sum int64
	
	for i := range fib_array {
		if i == 0 {
			fib_array[0] = 0
		} else if i == 1 {
			fib_array[1] = 1
		} else {
			fib_array[i] = fib_array[i - 1] + fib_array[i - 2]
		}
		//fmt.Println(fib_array)
		if (fib_array[i] % 2 == 0) && (fib_array[i] < 4000000) {
			sum += fib_array[i]
		}
		
	}
	
	fmt.Println(fib_array)
	return sum
}

func main() {
	
	found := fibonacci()
	fmt.Println(found)	
}
