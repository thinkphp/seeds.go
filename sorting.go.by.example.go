package main

import ("fmt"
        "sort")

func main() {

     strs := []string{"c", "a", "b"}
     ints := []int{9,8,7,6,5,4,3,2,0}

     sort.Strings(strs)
     sort.Ints(ints)

     i := sort.IntsAreSorted(ints)
     s := sort.StringsAreSorted(strs)

     fmt.Println(ints)
     fmt.Println(strs)

     fmt.Println(i)
     fmt.Println(s)

}