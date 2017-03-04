package main

import ("fmt"
       "os"
       "strconv") 

func main() {

     var sieve [60001]int

     var n int

     var j int 

     n , _ = strconv.Atoi(os.Args[1])

     for i:= 2; i * i <= n ; i++ {

         if sieve[ i ] == 0 {

            j = 2

            for i * j <= n {

                sieve[ i * j ] = 1

                j++
            }
         }
     }

     for i:= 2; i <= n ; i++ {

         if sieve[ i ] == 0 {

            fmt.Printf("%d ", i)
         } 
     }

}