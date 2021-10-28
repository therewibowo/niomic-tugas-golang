package main

import "fmt"
import "strconv"

func main()  {
	
	var input string
	fmt.Println("masukan angka : ")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)

	if err == nil{
		defer fmt.Println("woy") 
		fmt.Println(number, " ini adalah angka")
	}else{
		defer fmt.Println("woy")
		fmt.Println(input, " Bukan angka ini")
		fmt.Println(err.Error())
	}

}