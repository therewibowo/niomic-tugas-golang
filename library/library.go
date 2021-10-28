package library

import "fmt"


func Tampikan_mahasiswa()  {
var n1 mahasiswa
n1.nama= "yuli"
n1.umur = 10

n1.data_mahasiswa()
}

type mahasiswa struct {
	nama string
	umur int
}

func (m mahasiswa)data_mahasiswa()  {
	fmt.Println("nama saya : " ,m.nama)
	fmt.Println("umur : " ,m.umur)
}

