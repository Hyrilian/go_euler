package main

import (
	"fmt"
	)
	
func main() {
	
	var added int
	
	for i := 0; i < 1000; i++ {
		switch {
		case i % 3 == 0:
			added += i
			fmt.Println(i)
		case i % 5 == 0:
			added += i
			fmt.Println(i)
		}
		
	}
	
	fmt.Println(added)
}

