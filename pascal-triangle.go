/**
 *
 * Adrian Statescu (http://adrianstatescu.com) 
 *
 * Pascal's Triangle
 *
 * License MIT
 *
 */
package main

import ("fmt"
       "os"
       "strconv")  

func displayPascalTriangle(n int) {    

     var pascal [ 100 ][ 100 ]int

     var i, j, k, sp int   

     pascal[ 0 ][ 0 ] = 1
 
     for i = 1; i <= n; i++ {

         for j = 0; j <= i; j++ {

             if 0 == j || j == i {

                pascal[ i ][ j ] = 1

             } else {

                pascal[ i ][ j ] = pascal[ i - 1][ j ] + pascal[ i - 1 ][ j - 1 ]
             }
         }
     }

     sp = n

     for i = 0; i <= n; i++ {

             for k = sp; k >= 0; k-- {

                     fmt.Printf(" ")
             }

             sp--
 
         for j = 0; j <= i; j++ {

             fmt.Printf("%d ", pascal[i][j])
         }

             fmt.Printf("\n")
     }
                
}

func main() {

     var n int 

     n,_ = strconv.Atoi(os.Args[1])

     displayPascalTriangle( n )           
}