package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/test_string")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	for i := 0; i < 5e4; i++ {
		sqlStr := `INSERT INTO test_deep_page (a, b, c) VALUES `
		for j := 0; j < 100; j++ {
			a := rand.Intn(10)
			b := rand.Intn(10)
			c := rand.Intn(10)
			sqlStr += fmt.Sprintf("(%d,%d,%d), ", a, b, c)
		}
		//fmt.Println(sqlStr[:len(sqlStr)-2])
		sqlStr = sqlStr[:len(sqlStr)-2]
		//fmt.Println(sqlStr)
		_, err := db.Exec(sqlStr)
		if err != nil {
			log.Fatalln(err)
		}
	}

	tx, _ := db.Begin()
	tx.Commit()
}
