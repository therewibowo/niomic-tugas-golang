package main

import "fmt"
import "database/sql"
import _ "mysql"

type guru struct{
	NIP int
	nama_guru string
	jenis_kelamin string
	jabatan string
}

func main() {
	//tambah_data()
	//ubah_data()
	hapus_data()
	sql_tampil()

}

func koneksi()(*sql.DB, error)  {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1)/tugas_golang")

	if err != nil {
		return nil, err
	}
	return db, nil

}
//fungsi tampil data
func sql_tampil()  {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	rows, err := db.Query("select* from tbl_guru ")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []guru
	for rows.Next(){
		var each = guru{}
		var err = rows.Scan(&each.NIP, &each.nama_guru, 
			&each.jenis_kelamin, &each.jabatan)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result{
		fmt.Println(each.NIP, each.nama_guru)
	}
}
func hapus_data()  {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
 _, err = db.Exec("delete  from tbl_guru where nama_guru =?","R. Darul Ulum")	

 if err != nil {
	 fmt.Println(err.Error())
	 return
 }
 fmt.Println("data berhasil di hapus")	
}

/*
func ubah_data()  {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
 _, err = db.Exec("update tbl_guru set jabatan =? where nama_guru =?","guru mata pelajaran","R. Darul Ulum")	

 if err != nil {
	 fmt.Println(err.Error())
	 return
 }
 fmt.Println("data berhasil di ubah")
}
/*
/*
func tambah_data()  {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into tbl_guru values(?,?,?,?) ",191291,"R. Darul Ulum","laki-laki","Wakil Kurikulum")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("input data berhasil")

}
*/