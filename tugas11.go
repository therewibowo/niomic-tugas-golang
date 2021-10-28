package main

import "fmt"
import "strconv"

func main() {
	var bil1 = "29"
	var bil2 = 19
	var num,err = strconv.Atoi(bil1)

	

	if err == nil {
		fmt.Println("pertambahan : "  , num+bil2)
		fmt.Println("perkalian : " , num * bil2)
		fmt.Println("pengurangan : " , num - bil2)
		fmt.Println("pembagian : " , num / bil2)
	}
}