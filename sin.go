/**
 *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
 *
 *   Approximation of SIN(x) using Taylor Series
 *   x - x^3/3! + x^5/5! - x^7/7! + ...
 *
 *   Compile and Run -> go run sin.go
 *
 *   License: MIT
 */

package main

import ("fmt"
        "os"
        "strconv" 
        "math")

func fact(x float64) float64 {

     if x == 0 { 

        return 1
     }

     return x * fact( x - 1 )
}

func pow(x,y float64) float64 {

     var p,i float64
       
     p = 1

     for i = 1; i <= y; i++ {

         p = p * x
     }

   return p
}

func sin(x float64) float64 {

    var t1, t2, r, i float64
 
     t1 = x

     t2 = t1 - float64(pow(x, 3)/fact(3))

     eps := 0.00001

     if t1 > t2 {

          r = t1 - t2 

     }  else {

          r= t2 - t1
     }

     i = 2

     for r >= eps {

          t1 = t2

          t2 += pow(-1,i) * float64(pow(x, 2*i+1)/fact(2*i+1))


          if t2 > t1 {

              r = t2 - t1 

          }  else {

              r= t1 - t2    
          }

          i += 1
     }         

    return t2
} 

func main() {

     var x float64

     x,_ = strconv.ParseFloat(os.Args[1], 64)
 
     fmt.Printf("sin(%.1f) %.50f\n", x, sin(x))

     fmt.Printf("sin(%.1f) %.50f", x, math.Sin(x))
}