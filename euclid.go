/**
 * Euclid's algorithm is an efficient method for computing the greatest common divisor of two numbers, I mean the largest number that divides both of them without leaving a remainder. 
 */
package main

import "fmt"

func gccRec(x, y int) int {

         if y == 0 {

            return x
         } 
         return gccRec(y, x % y)
}

func gccIte(a, b int) int {    

     var r int

     for b > 0 {

         r = a % b

         a = b

         b = r  
     } 

     return a
}

func main() {

     fmt.Println(gccRec(10,30));  

     fmt.Println(gccIte(10,3));  
}