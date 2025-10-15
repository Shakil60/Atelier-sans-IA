package main

import "ateliersansia"
import "fmt"

func main(){
    fmt.Println(ateliersansia.CompressRLE("aaabbc")) // a3b2c1
    fmt.Println(ateliersansia.CompressRLE("ab"))     // a1b1
}