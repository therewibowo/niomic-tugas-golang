package main

import "fmt"

func main()  {
	var buah = []string{"apel","jeruk","anggur","melon"}
	buah = append(buah,"durian")
	fmt.Println("Jumlah Element :" ,len(buah))

	fmt.Println("isi Element : " ,(buah))
		
	for i := 0; i < 5; i++ {
		
		fmt.Println("Element ke - : ", i ,(buah[i]))
		
		} 
}