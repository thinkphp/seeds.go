package main

import ("fmt"
       "os"
       "strconv")

func generateSubsets(n int) {

     arr := make([]int, n) 

     var i, s int

     for s < n {

         arr[ n - 1 ]++
  
         for i := len(arr) - 1; i > 0; i-- {

             if arr[i] > 1 {
       
                arr[i] -= 2

                arr[i - 1] += 1   
             }
         }    

         for i = 0; i <= n - 1; i++ {

             if arr[ i ] != 0 {

                fmt.Printf("%d ", (i + 1)) 
             }
         }

         s = 0

         for i = 0; i <= n - 1; i++ {

             if arr[ i ] == 1 {

                s++  
             }
         }

         fmt.Printf("\n") 
                            
     }      
}


func main() {

     var n,_ = strconv.Atoi(os.Args[1])

     generateSubsets( n )
}
