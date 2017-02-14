package main

import ("fmt"
       "math/rand")

func bubblesort(arr []int) {

     done := true
 
     var changes, i int 
 
     for done {

         changes = 0
         
         for i = 0; i < len(arr) - 1; i++ {

             if arr[ i ] > arr[ i + 1 ] {

                arr[ i ], arr[ i + 1 ] = arr[ i + 1 ], arr[ i ]

                changes++ 
             }   
         } 

         if changes == 0 {

            done = false
         }
     }
}

func main() {

     size := 100

     var arr = make([]int, size)

     index := 0

     for index < size {

         arr[ index ] = rand.Intn(size)

         index++ 
     } 

     bubblesort( arr )

     for _, v := range arr {

         fmt.Printf("%d ", v)
     }   
}