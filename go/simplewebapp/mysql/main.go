package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "testusr:password@/testdb")
	checkErr(err)

	// CREATE TABLE `userinfo` (
	//     `uid` INT(10) NOT NULL AUTO_INCREMENT,
	//     `username` VARCHAR(64) NULL DEFAULT NULL,
	//     `departname` VARCHAR(64) NULL DEFAULT NULL,
	//     `created` DATE NULL DEFAULT NULL,
	//     PRIMARY KEY (`uid`)
	// )
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("testusr", "No8", time.Now().Unix())
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("affect:", affect)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id:", id)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created int
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid, username, departname, created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
	fmt.Println("END")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
