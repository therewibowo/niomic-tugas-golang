package main

import "fmt"

func main()  {
	tampilkan_mahasiswa()
}

func tampilkan_mahasiswa()  {
	var nama_mahasiswa = map[string]int{"aldo" : 182,"yosep": 178}
	fmt.Println("Aldo",nama_mahasiswa["aldo"],"cm")
	fmt.Println("Yosep",nama_mahasiswa["yosep"],"cm")
}