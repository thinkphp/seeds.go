/*   
   Author     : Adrian Statescu @thinkphp http://adrianstatescu.com http://thinkphp.ro
   Description: Prime Factorization in Golang.
   10  = 2^1 * 5^1
   100 = 2^2 * 5^2
   15  = 3^1 * 5^1 
   MIT License
 */
package main

import ("fmt"
       "os"
       "strconv")

func main() {

     var n,
         i,
         multiply int

     n , _ = strconv.Atoi(os.Args[1])

     i = 2
 
     for n != 1 {

         multiply = 0

         for n % i == 0 {

             multiply++

             n /= i  
         }  

         if multiply != 0 {

            fmt.Printf("%d^%d\n", i, multiply)  
         }

         i++
     }        

}