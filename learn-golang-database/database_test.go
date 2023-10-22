package learn_golang_database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

const dataSource = "root:root@tcp(localhost:3306)/learn_golang_database"

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	fmt.Println("Koneksi Database Sukses")
	db.Close()
}
