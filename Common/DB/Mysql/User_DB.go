package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

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
