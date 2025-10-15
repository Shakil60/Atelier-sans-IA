package main

import "ateliersansia"
import "fmt"

func main(){
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(ateliersansia.MaxSubArray(nums)) // 6 (le sous-tableau [4, -1, 2, 1])
}