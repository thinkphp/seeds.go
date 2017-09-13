package main

import "fmt"

func GDC(a, b int) int{

     for b != 0 {

         a, b = b, a % b
     }
 
     return a
}

func _nEuclid(lo, hi int, vec []int) int {

     if lo == hi {

        return vec[ lo ] 
     }      

     var mi = (lo + hi)>>1

     return GDC(_nEuclid(lo, mi, vec), _nEuclid(mi + 1, hi, vec)) 
}

func nEuclid(arr []int) int {

     return _nEuclid(0, len(arr) - 1, arr)  
}

func main() {

     arr := []int{100, 40, 80}

     fmt.Println(nEuclid(arr))
}