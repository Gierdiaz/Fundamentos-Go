//Deslocamento de bits
package main

import ("fmt")

const  
( 
	_ = iota 			  // 0 << (0*10)
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 2 << (2*20)
	GB = 1 << (iota * 10) // 3 << (3*30)
	TB = 1 << (iota * 10) // 4 << (4*40)
)

func main () {
	fmt.Println("Binary\t\t\t\tDecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}