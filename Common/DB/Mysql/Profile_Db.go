package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insertprofile(Profileresponse Model.Profile) (Profilesend Model.Profile) {

	fmt.Println("inside Database")
	var Count int

	query := "SELECT COUNT(*) FROM profile WHERE profile.loginid ='" + Profileresponse.Loginid + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		row, err := OpenConnection["Rentmatics"].Exec("insert into profile (loginid,name,tittle,phone,email,about,twitter,facebook,google) values (?,?,?,?,?,?,?,?,?)", Profileresponse.Loginid, Profileresponse.Name, Profileresponse.Tittle, Profileresponse.Phone, Profileresponse.Email, Profileresponse.About, Profileresponse.Twiter, Profileresponse.Facebook, Profileresponse.Google)
		fmt.Println("successfully inserted", row, err)
		rows, err := OpenConnection["Rentmatics"].Query("select * from profile  where loginid=?", Profileresponse.Loginid)
		for rows.Next() {

			rows.Scan(

				&Profilesend.Id,
				&Profilesend.Loginid,
				&Profilesend.Name,
				&Profilesend.Tittle,
				&Profilesend.Phone,
				&Profilesend.Email,
				&Profilesend.About,
				&Profilesend.Twiter,
				&Profilesend.Facebook,
				&Profilesend.Google,
			)

		}

		return Profilesend
	} else {
		Queryupdate := "UPDATE profile SET name= " + Profileresponse.Name + " ,tittle=" + Profileresponse.Tittle + " ,phone= '" + Profileresponse.Phone + "' ,email= '" + Profileresponse.Email + "' ,about= '" + Profileresponse.About + "' ,twitter= '" + Profileresponse.Twiter + "'  ,facebook= '" + Profileresponse.Facebook + "'  ,google= " + Profileresponse.Google + "'  where loginid= " + Profileresponse.Loginid
		row, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
		fmt.Println("successfully inserted", row, err)

		rows, err := OpenConnection["Rentmatics"].Query("select * from profile  where loginid=?", Profileresponse.Loginid)

		for rows.Next() {

			rows.Scan(

				&Profilesend.Id,
				&Profilesend.Loginid,
				&Profilesend.Name,
				&Profilesend.Tittle,
				&Profilesend.Phone,
				&Profilesend.Email,
				&Profilesend.About,
				&Profilesend.Twiter,
				&Profilesend.Facebook,
				&Profilesend.Google,
			)

		}

		return Profilesend

	}
}

func Getprofile(Profileresponse Model.Getprofile) (Profilesend Model.Profile) {

	rows, _ := OpenConnection["Rentmatics"].Query("select * from profile  where loginid=?", Profileresponse.Loginid)
	for rows.Next() {

		rows.Scan(

			&Profilesend.Id,
			&Profilesend.Loginid,
			&Profilesend.Name,
			&Profilesend.Tittle,
			&Profilesend.Phone,
			&Profilesend.Email,
			&Profilesend.About,
			&Profilesend.Twiter,
			&Profilesend.Facebook,
			&Profilesend.Google,
		)

	}

	return Profilesend
}
