package Mysql

import (
	_ "database/sql"

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
		if err != nil {
			log.Error("Error -DB: All Home", err, row)
		}

	} else {

		log.Info("Already Exist")

	}
}
