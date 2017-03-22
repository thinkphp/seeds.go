package main 

import ("fmt"
        "os"
        "strconv" 
        "math")

func f(x, a float64) float64 {
 
     return math.Exp(x) - a
}

func ln(n float64) float64 {

     var lo, hi, m float64     

     if n <= 0 {

        return -1
     }

     if n == 1 {

        return 0 
     }

     EPS := 0.00001  

     lo = 0

     hi = n

     for math.Abs(lo - hi) >= EPS {

         m = float64( (lo + hi) / 2.0 )

         if f(m, n) < 0 {

              lo = m;

         } else {

              hi = m  
         }         
     }

  return float64((lo+hi)/2.0)
}

func main() {

     var n float64 

     n,_ = strconv.ParseFloat(os.Args[1], 64)

     fmt.Printf("ln(%f) = %.19f\n", n, ln( n ))    

     fmt.Printf("ln(%f) = %.19f", n, math.Log( n )) 