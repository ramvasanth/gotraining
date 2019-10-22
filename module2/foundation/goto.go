package main

import "fmt"

func labelFunc() {
out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i+j == 20 {
				break out
			}
		}
	}
}

func labelFunc2() {
	if 1 == 1 {
		fmt.Println(1)
		goto End
	}
	fmt.Println(2)
End:
	fmt.Println(3)
}
