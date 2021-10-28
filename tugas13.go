package main

import "fmt"
import "encoding/base64"
import "crypto/sha1"

func main() {
	var nama = "Yuli Prihati, S.Kom"

	//proses enkripsi sha1
	var kodesha1 = sha1.New()
	kodesha1.Write([]byte(nama))

	var enkrip = kodesha1.Sum(nil)
	var stringenkrip = fmt.Sprintf("%x", enkrip)

	

	//proses enkripsi
	var encodestring = base64.StdEncoding.EncodeToString([]byte(nama))
	fmt.Println("encode string ", encodestring)

	//proses dekripsi
	var decodeByte,_ = base64.StdEncoding.DecodeString("WXVsaSBQcmloYXRpLCBTLktvbQ==")
	var decodestring = string(decodeByte)
	fmt.Println("isi dari string adalah : ", decodestring)

	fmt.Println("enkripsi dari : ",nama, stringenkrip)
}