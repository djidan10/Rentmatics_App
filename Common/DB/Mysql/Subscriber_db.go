package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Insertsubscribtion(Subuser Model.Subscribeuser) {

	row, err := OpenConnection["Rentmatics"].Exec("insert into subscriber (loginid,city) values (?,?)", Subuser.Sub_Email, Subuser.Sub_City)
	if err != nil {
		log.Error("Error -DB: All Home", err, row)

	}
}
