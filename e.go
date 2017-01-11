/**
 *
 * Adrian Statescu (http://adrianstatescu.com) 
 *
 * The number e is a mathematical constant that is the base of the natural logarithm.
 * Sometimes called Euler's number.
 *
 * License MIT
 *
 */
package main

import "fmt"

func fact(n int) int {    

     if n == 0 {

        return 1

     } else {

       return n * fact(n - 1)
     }
}

func e(eps float64) float64 {

     var term1, term2, r float64

     var i int 

     term1 = 1.0       
 
     term2 = term1 + float64(1.0/fact(1)) 

     i = 2

     r = term2 - term1 

     for r > eps {

          term1 = term2

          term2 += 1.0/float64(fact(i))

          r = term2 - term1 

          i += 1
     }  

     return term2
}

func main() {

     fmt.Printf("%.30f", e(0.00001))
}