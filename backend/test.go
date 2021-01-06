package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type TODO struct {
	Limit     string
	Title     string
	Intro     string
	Share     bool
	Islimited bool
}

func input() {
	db, err := sql.Open("postgres", "user=postgres password=seriver25859 dbname=todo sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO todos VALUES ($1,$2,$3,$4,$5);")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec("2020-1-3", "test2", "testTask2", 1, 1)
	if err != nil {
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func output() {
	db, err := sql.Open("postgres", "user=postgres password=seriver25859 dbname=todo sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		t := TODO{}
		rows.Scan(&t.Limit, &t.Title, &t.Intro, &t.Share, &t.Islimited)
		fmt.Println(t)
	}
}

func main() {
	// input()
	output()

}
