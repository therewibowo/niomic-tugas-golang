package main

import "fmt"
import "reflect"
import "runtime"

func main()  {

	runtime.GOMAXPROCS(2)
	
	var input string
	fmt.Scanln(&input)
	


 go tampilkan_pesan()

 tampilkan_pesan()

}

func tampilkan_pesan()  {
	var number = 10
	var nama = "Tri Wibowo"

	var reflect_number = reflect.ValueOf(number);
	var reflect_nama = reflect.ValueOf(nama)
	fmt.Println("tipe variable: ", reflect_number.Type())
	fmt.Println("tipe variabel nama : ", reflect_nama.Type())
	
}

// func tampilkan_pesan(x int, pesan string)  {
// 	for i := 0; i < x; i++ {
// 		fmt.Println((i+1), pesan)
// 	}	
// }
