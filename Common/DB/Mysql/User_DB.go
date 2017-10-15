package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	"net/smtp"

	_ "github.com/go-sql-driver/mysql"
)

func InserUser(Userdata Model.User) (Userinfo Model.UserResponse) {
	var Count int

	query := "SELECT COUNT(*) FROM userdata WHERE userdata.loginid ='" + Userdata.Loginid + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		row, err := OpenConnection["Rentmatics"].Exec("insert into userdata (username,password,loginid,login_type) values (?,?,?,?)", Userdata.Username, Userdata.Password, Userdata.Loginid, Userdata.Logintype)
		if err != nil {
			log.Error("Error -DB: User", err, row)
		}
		rows, err := OpenConnection["Rentmatics"].Query("select username,loginid from userdata where loginid=?", Userdata.Loginid)
		if err != nil {
			log.Error("Error -DB: User", err)
		}
		for rows.Next() {

			rows.Scan(

				&Userinfo.Username,
				&Userinfo.Loginid,
			)

		}
		Userinfo.Status = "Success"
		return Userinfo
	} else {

		rows, _ := OpenConnection["Rentmatics"].Query("select username,loginid from userdata where loginid=?", Userdata.Loginid)
		for rows.Next() {

			rows.Scan(
				&Userinfo.Username,
				&Userinfo.Loginid,
			)

		}

		Userinfo.Status = "Already_Exist"
		return Userinfo

	}
}

//Rentmatics Admin Login
func InserRentUser(Userdata Model.RentUser) (Status Model.AdminResponse) {
	var Count int

	query := "SELECT COUNT(*) FROM userdata WHERE userdata.username ='" + Userdata.Username + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		row, err := OpenConnection["Rentmatics"].Exec("insert into rentmatics_userdetails (username,password,role) values (?,?,?)", Userdata.Username, Userdata.Password, Userdata.Role)
		if err != nil {
			log.Error("Error -DB: User", err, row)
		}
		rows, err := OpenConnection["Rentmatics"].Query("select username,role from rentmatics_userdetails where username=?", Userdata.Username)
		if err != nil {
			log.Error("Error -DB: User", err)
		}
		for rows.Next() {

			rows.Scan(
				&Status.Username,
				&Status.Role,
			)

		}

		return
	} else {

		rows, err := OpenConnection["Rentmatics"].Query("select username,role from rentmatics_userdetails where username=?", Userdata.Username)
		if err != nil {
			log.Error("Error -DB: User", err)
		}
		for rows.Next() {

			rows.Scan(
				&Status.Username,
				&Status.Role,
			)

		}
		return
	}

}

func GetUSer(User1 Model.LoginUser) (Userinfo Model.UserResponse) {
	var Getuser Model.Getlogin
	rows, err := OpenConnection["Rentmatics"].Query("select username,loginid,password from userdata where username=?", User1.Username)
	if err != nil {
		log.Error("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Getuser.Username,
			&Getuser.Loginid,
			&Getuser.Password,
		)

	}
	if User1.Password == Getuser.Password {

		Userinfo.Username = Getuser.Username
		Userinfo.Loginid = Getuser.Loginid
		Userinfo.Status = "Success"
		return

	} else {
		Userinfo.Status = "Failure"
		return
	}

}

func GetRentUSer(User1 Model.LoginUser) (Adminres Model.AdminLoginResponse) {
	rows, err := OpenConnection["Rentmatics"].Query("select username,password,role from rentmatics_userdetails where username=?", User1.Username)
	if err != nil {
		log.Error("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(
			&Adminres.Username,
			&Adminres.Password,
			&Adminres.Role,
		)

	}
	if User1.Password == Adminres.Password {

		return Adminres

	} else {
		Adminres.Role = "Invalid"
		return Adminres
	}

}

func InsertChangepassword(User1 Model.Changepass) (Userinfo string) {
	var Getuser string
	rows, err := OpenConnection["Rentmatics"].Query("select password from userdata where loginid=?", User1.Loginid)
	if err != nil {
		log.Error("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Getuser,
		)

	}
	fmt.Println(Getuser)
	fmt.Println(User1.OldPassword)
	if User1.OldPassword == Getuser {

		Queryupdate := "UPDATE  userdata SET password='" + User1.Newpassword + "'where loginid='" + User1.Loginid + "'"

		row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
		if err != nil {
			log.Error("Error -DB: update Profile", err, row)
		}
		return "success"

	} else {

		return "Failure"
	}

}

func InsertForgotspassword(User1 string) string {
	var Username string
	var Password string
	var Count1 int

	query := "SELECT COUNT(*) FROM userdata WHERE userdata.loginid ='" + User1 + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count1,
	)
	if Count1 != 0 {

		rows, err := OpenConnection["Rentmatics"].Query("select username,password from userdata where loginid=?", User1)
		if err != nil {
			log.Error("Error -DB: Get User", err)
		}
		for rows.Next() {

			rows.Scan(

				&Username,
				&Password,
			)

		}

		auth := smtp.PlainAuth("", "Rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
		var receiptent = "To:" + User1 + "\r\n"
		to := []string{User1}
		msg := []byte(receiptent +
			"Subject: RENTMATICS NOTIFICATION!\r\n" +
			"\r\nUsername:" + Username + "\r\nPassword:" + Password +
			"\r\n")

		mailerr := smtp.SendMail("smtp.gmail.com:587", auth, "Rentmatics@gmail.com", to, msg)
		if mailerr != nil {
			log.Error("Mail Error - Notification", mailerr)

		}
		return "Success"
	} else {
		return "Failure"

	}

}
