package ateliersansia

import "fmt"

func PrintCenteredPyramid(n int){
	for i := 0; i <= n; i++ {
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
	fmt.Print("\n")
	} 
}