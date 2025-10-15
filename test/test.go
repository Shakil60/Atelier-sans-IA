package main

import "ateliersansia"
import "fmt"

func main(){
    nums := []int{2, 7, 11, 15}
    target := 9
    i, j := ateliersansia.TwoSum(nums, target)
    fmt.Println(i, j) // Résultat attendu : 0 1
}