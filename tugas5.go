package main

import "fmt"

type pelajar struct {
 nama string
 umur int
}
func main(){
var n1,n2,n3 pelajar;
n1.nama = "Ahmad"
n1.umur = 81
n2.nama = "Abi"
n2.umur = 87
n3.nama = "Zaka"
n3.umur = 47

n1.datasaya()
n2.datasaya()
n3.datasaya()

}

func (n pelajar)datasaya()  {
	fmt.Println(n.nama, n.umur)
}
