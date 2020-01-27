package main

import ("fmt"
	    "math")

func quadratic(a, b, c float64) (x1, x2 float64) {

     bSquared := b * b

     fourAC := 4 * a * c

     twoA := 2 * a

     D := math.Sqrt(bSquared - fourAC)     

     if D > 0 {

        x1  = (- b + D) / twoA

        x2  = (- b - D) / twoA

        x1 =fourAC

        x2 = D 

     } else if D == 0 {

     	x1, x2 = - b / 2 * a, - b / 2 * a

     } else {

        x1, x2 = -1, -1
     }     

     return
}	    

func degreeOne(a, b float64) (x float64) {

	 if (a == 0 && b == 0) {

	 	fmt.Println("Equation has infinity solutions!")

	 } else if (a == 0 && b != 0) {

        fmt.Println("Equation has not solutions!")	 	

	 } else if (a != 0 && b != 0) {

	         x = - b / a	

	 } else if (a != 0 && b == 0) {

            x = 0
	 }
 
	 return
}

func main() {

   x1, x2 := quadratic(1, 2, -2)

   fmt.Println(x1, x2)
   x := degreeOne(2, 3)
   if x == 0 || x != 0 { fmt.Println(x) }

}
