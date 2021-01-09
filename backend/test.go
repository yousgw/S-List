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

// 更新時の参照用としてTaskidとか作ったほうがいい

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func in(t TODO) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO todos VALUES ($1,$2,$3,$4,$5);")
	checkErr(err)
	res, err := stmt.Exec(t.Limit, t.Title, t.Intro, t.Share, t.Islimited)
	checkErr(err)
	_ = res
}

// 期日順で出力
func outAll() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM todos order by limit_date")
	checkErr(err)
	for rows.Next() {
		t := TODO{}
		rows.Scan(&t.Limit, &t.Title, &t.Intro, &t.Share, &t.Islimited)
		fmt.Println(t)
	}
}

// 期日ありのタスクを期日順で出力
func outLimited() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM todos WHERE islimited='t' order by limit_date")
	checkErr(err)
	for rows.Next() {
		t := TODO{}
		rows.Scan(&t.Limit, &t.Title, &t.Intro, &t.Share, &t.Islimited)
		fmt.Println(t)
	}
}

// 期日なしのタスクを期日(作成日)順で出力
func outUnlimited() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM todos WHERE islimited='f' order by limit_date")
	checkErr(err)
	for rows.Next() {
		t := TODO{}
		rows.Scan(&t.Limit, &t.Title, &t.Intro, &t.Share, &t.Islimited)
		fmt.Println(t)
	}
}

func main() {
	// in()
	outUnlimited()

}
