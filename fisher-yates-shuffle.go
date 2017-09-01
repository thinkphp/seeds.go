package main

import "fmt"
import "math/rand"
import "time"

func main() {

    fmt.Println("Fisher-Yates Shuffle Algorithm ")

    i := 0

    s1 := rand.NewSource(time.Now().UnixNano())

    arr := [10]int{1,2,3,4,5,6,7,8,9,0}

    n := 10

    fmt.Println(arr)

    for i < n {

       r1 := rand.New(s1)

       R := r1.Intn(i+1)

       arr[i], arr[R] = arr[R], arr[i]
              
       i += 1 
    }  

    fmt.Println(arr)
}