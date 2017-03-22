/**
 *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
 *
 *   Approximation of COS(x) using Taylor Series
 *   1 - x^2/2! + x^4/4! - x^6/6! + ...
 *
 *   Compile and Run -> go run cos.go 2
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

func cos(x float64) float64 {

    var t1, t2, r, i float64
 
     t1 = 1

     t2 = t1 - float64(pow(x, 2)/fact(2))

     eps := 0.00001

     if t1 > t2 {

          r = t1 - t2 

     }  else {

          r= t2 - t1
     }

     i = 2

     for r >= eps {

          t1 = t2

          t2 += pow(-1,i) * float64(pow(x, 2*i)/fact(2*i))


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
 
     fmt.Printf("cos(%.1f) %.50f\n", x, cos(x))

     fmt.Printf("cos(%.1f) %.50f", x, math.Cos(x))
}