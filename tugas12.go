package main

import "fmt"
import "regexp"

func main() {
	var nama = "asus samsung19 acer "
	var regex, err = regexp.Compile(`[A-Z]+\d `)

	if err != nil {
		fmt.Println(err.Error())
	}

	var nama_laptop = regex.ReplaceAllString(nama,"SAMSUNG")
	fmt.Println(nama_laptop)

	var cocok = regex.MatchString(nama)
	fmt.Println(cocok)
}