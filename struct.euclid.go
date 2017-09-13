package main

import "fmt"

type Pair struct {
     a, b int
}

func (n *Pair) euclid() int {
      
      for n.b != 0 {

          n.a, n.b = n.b, n.a % n.b
      }

      return n.a
}

func nEuclid(lo, hi int, arr []int) int {

     if lo == hi {

        return arr[lo] 
     }  

     mi := (lo+hi)>>1

     r := Pair{nEuclid(lo, mi, arr), nEuclid(mi+1, hi, arr)} 

     return r.euclid()
}


func main() {

     a := 10
     b := 4
     c := 10 
     d := 3

     r := Pair{a,b}
     f := Pair{c,d}

     fmt.Printf("GDC<%d,%d> = %d\n", a, b, r.euclid()) 
     fmt.Printf("GDC<%d,%d> = %d\n", c, d, f.euclid()) 

     arr := []int{100,3,20,30,40,50}
     fmt.Println(nEuclid(0, len(arr) - 1, arr))

}

