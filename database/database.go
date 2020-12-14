package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("insert into userinfo set username=? ,department=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "haha")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Pritln(uid, username, department, created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid = ?")
	checkErr(err)
	res, _ = stmt.Exec(id)
	after, _ = res.RowsAffected()
	fmt.Println(after)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
