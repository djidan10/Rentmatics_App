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

func InsertBookings(book Model.Booknow) string {

	row, err := OpenConnection["Rentmatics"].Exec("insert into  Rentbookings (Homeid,Loginid,Amount,Buyer_name,Email,Phone,Paymentid,Refercode,Description) values (?,?,?,?,?,?,?,?,?)", book.Homeid, book.Loginid, book.Amount, book.Buyer_name, book.Email, book.Phone, book.Paymentid, book.Referalcode, book.Description)
	if err != nil {
		log.Error("Error -DB: All Home", err, row)

	}

	return "success"
}
