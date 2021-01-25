package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)

type TODO struct {
	Limit     string `json:limit`
	Title     string `json:title`
	Intro     string `json:intro`
	Share     bool `json:isShared`
	Islimited bool `json:isLimited`
}

// 更新時の参照用としてTaskidとか作ったほうがいい

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func In(c *gin.Context) {
	var t TODO
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	c.BindJSON(&t)

	stmt, err := db.Prepare("INSERT INTO todos VALUES ($1,$2,$3,$4,$5);")
	checkErr(err)
	res, err := stmt.Exec(t.Limit, t.Title, t.Intro, t.Share, t.Islimited)
	checkErr(err)
	_ = res
	c.String(http.StatusCreated, "OK")
}

// 期日順で出力
func OutAll(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo sslmode=disable")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM todos order by limit_date")
	checkErr(err)
	var Data []TODO
	for rows.Next() {
		t := TODO{}
		rows.Scan(&t.Limit, &t.Title, &t.Intro, &t.Share, &t.Islimited)
		fmt.Println(t)
		Data = append(Data, t)
	}
	c.JSON(http.StatusOK, Data)
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

func Test(c *gin.Context){
	c.JSON(http.StatusOK,"OK")
}
