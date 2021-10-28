package main

import "fmt"
import "os/exec"

func main() {
	//kalkulator
	var hasil,_ = exec.Command("calc").Output()
	var explor,_ = exec.Command("explorer").Output()
	var dxdiag,_ = exec.Command("dxdiag").Output()
	fmt.Println(string(hasil))
	fmt.Println(string(explor))
	fmt.Println(string(dxdiag))
}