package main

import "fmt"

func euclid_iterative(a, b int) int {

     var r int

     for b != 0 {

         r = a % b

         a = b

         b = r
     } 

     return a
}

func euclid_iterative2(a, b int) int {

     var tmp int

     for b != 0 {

         tmp = b

         b = a % b 

         a = tmp
     } 

     return a
}


func euclid_recursive(a, b int) int {

     if b == 0 { 

        return a

     } else {

        return euclid_recursive(b, a % b)  
     }
}

func euclid_substraction(a, b int) int {

     for a != b {

         if a > b {

            a -= b  

         } else {

            b -= a
         }
     } 

     return a 
}


func main() {

     fmt.Printf("GDC[%d, %d] = %d; Should be 1\n", 10, 3, euclid_iterative(10,3))
     fmt.Printf("GDC[%d, %d] = %d; Should be 1\n", 10, 3, euclid_iterative2(10,3))
     fmt.Printf("GDC[%d, %d] = %d; Should be 1\n", 10, 3, euclid_recursive(10,3))
     fmt.Printf("GDC[%d, %d] = %d; Should be 1\n", 10, 3, euclid_substraction(10,3))
}