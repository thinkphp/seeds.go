/**
 *
 * Adrian Statescu (http://adrianstatescu.com) 
 *
 * Comb using Pascal's Triangle
 *
 * License MIT
 *
 */
package main

import ("fmt"
       "os"
       "strconv")  

func comb(n int, k int) int {    

     var pascal [ 500 ][ 500 ]int

     var i, j int   

     pascal[ 0 ][ 0 ] = 1
 
     for i = 1; i <= n; i++ {

         for j = 0; j <= i; j++ {

             if 0 == j || j == i {

                pascal[ i ][ j ] = 1

             } else {

                pascal[ i ][ j ] = pascal[ i - 1 ][ j ] + pascal[ i - 1 ][ j - 1 ]
             }
         }
     }

     return pascal[ n ][ k ]                
}

func main() {

     var n, k int 

     n,_ = strconv.Atoi(os.Args[1])

     k,_ = strconv.Atoi(os.Args[2])

     fmt.Printf("Comb(%d, %d) = %d", n, k,comb(n, k))
}