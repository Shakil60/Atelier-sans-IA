package main

import "ateliersansia"
import "fmt"

func main(){
    fmt.Println(ateliersansia.CompressRLE("aaabbccdde")) // "a3b2c2d2e1"
    fmt.Println(ateliersansia.CompressRLE("hellooo"))   // "h1e1l2o3"
}