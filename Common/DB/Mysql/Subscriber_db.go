package Mysql

import (
	//	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	//	"strconv"

	//"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Insertsubscribtion(Subuser string) {
	var Count int
	query := "SELECT COUNT(*) FROM subscriber WHERE subscriber.loginid ='" + Subuser + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		row, err := OpenConnection["Rentmatics"].Exec("insert into subscriber (loginid) values (?,)", Subuser)
		fmt.Println("successfully inserted", row, err)

	} else {

		fmt.Println("Already Exist")

	}
}
