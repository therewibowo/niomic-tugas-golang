package main

import "fmt"
import "net/http"

func index(web http.ResponseWriter, r*http.Request)  {
	for i := 1; i < 101; i++ {
		fmt.Fprintln(web,"nomor",i)
	
	}
}

func main() {
	http.HandleFunc("/",func (web http.ResponseWriter, r*http.Request)  {
		fmt.Fprintln(web,"halaman utama")
		
	})	
	http.HandleFunc("/index", index)
		
	fmt.Println("server berjalan di localhost:8000")
		http.ListenAndServe(":8000",nil)
		
}