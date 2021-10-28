package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql"
	"net/http"
)

type data_guru struct {
	NIP           int
	nama_guru     string
	jenis_kelamin string
	jabatan       string
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/tugas_golang ")

	if err != nil {
		return nil, err
	}
	return db, nil
}

var data []data_guru

func main() {
	ambil_data()
	http.HandleFunc("guru", ambil_guru)
	fmt.Println("menjalankan web server pada localhost 8080")
	http.ListenAndServe(":8080", nil)
}

func ambil_guru(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func ambil_data() {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tbl_guru")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = data_guru{}
		var err = rows.Scan(&each.NIP, &each.nama_guru, &each.jenis_kelamin, &each.jabatan)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
