/**
 *   Adrian Statescu <mergesortv@gmail.com> (http://adrianstatescu.com) 
 *
 *   Golang Program to compute PI number using Monte Carlo Method.
 *   Compile and Run -> go run monte_carlo_pi.go 9000000
 *
 *   License: MIT
 */
package main

import ("fmt" 
        "math/rand" 
        "os" 
        "strconv" 
        "time")
 
func main() {
      
     var n int

     n , _ = strconv.Atoi(os.Args[1])

     count := 0

     rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

     for i := 0; i < n; i++ {

       x := rnd.Float64()

       y := rnd.Float64() 
    
       z := x*x + y*y

       if( z <= 1) { 

           count++  
       }
     }

     pi := float64(count) / float64(n) * float64(4)

     fmt.Printf("%s", pi)
}