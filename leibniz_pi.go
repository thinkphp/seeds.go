/**
 *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
 *
 *   Approximation of PI using Leibniz method.
 *   1-1/3+1/5-1/7
 *
 *   Compile and Run -> go run leibniz_pi.go
 *
 *   License: MIT
 */

package main

import "fmt"

func main() {

     var t1, t2, r, sign, i float64
 
     t1 = 1.0

     t2 = 1.0 - float64(1.0/3.0)

     r = t1 - t2; 

     i = 5;

     eps := 0.00001

     sign = 1

     for float64(4 * r) >= eps {

          t1 = t2

          t2 += float64( 1.0 / i) * sign

          if t1 > t2 {
             
              r = float64(t1 - t2) 

          }  else {

              r = float64(t2 - t1)
          }

          sign = (-1)*sign

          i += 2
     }         

     fmt.Printf("%.50f", 4 * t2)           
}