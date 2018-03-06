package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	//"net/smtp"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func InserUser(Userdata Model.User) (Userinfo Model.UserResponse) {
	fmt.Println("inside rent insert")

	var Count int

	query := "SELECT COUNT(*) FROM userdata WHERE userdata.loginid ='" + Userdata.Loginid + "'"
	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		rand.Seed(time.Now().UnixNano())
		randomno := randomInt(1234567, 7654321)
		referid := "RENT_" + strconv.Itoa(randomno)
		fmt.Println("rent id", referid)

		row, err := OpenConnection["Rentmatics"].Exec("insert into userdata (username,password,loginid,login_type,	Referid) values (?,?,?,?,?)", Userdata.Username, Userdata.Password, Userdata.Loginid, Userdata.Logintype, referid)
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
	rows, err := OpenConnection["Rentmatics"].Query("select username,loginid,password from userdata where loginid=?", User1.Username)
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
	fmt.Println(Adminres)
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
		Tostring := "Username:" + Username + "<br>Password:" + Password

		sendkey := os.Getenv("SENDGRID_API_KEYGO")
		fmt.Println(sendkey)

		from := mail.NewEmail("Rentmatics User", "sandhiyabalakrishnan6@gmail.com")
		subject := "RENTMATICS NOTIFICATION - FORGOT PASSWORD!"
		to := mail.NewEmail("Example User", User1)
		plainTextContent := Tostring
		htmlContent := "<strong>RENTMATICS NOTIFICATION - FORGOT PASSWORD <br> <br></strong><p>" + Tostring + "</p><br><b >Thank you for Contacting Rentmatics</b>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		fmt.Println(User1)
		client := sendgrid.NewSendClient(sendkey)
		response, err := client.Send(message)
		if err != nil {
			log.Error(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)

		}

		return "Success"
	} else {
		return "Failure"

	}

}
func Insertfeedback(Insfeedback Model.Feedback) string {
	row, err := OpenConnection["Rentmatics"].Exec("insert into feedback (name,Emailid,rating) values (?,?,?)", Insfeedback.Name, Insfeedback.Emailid, Insfeedback.Rating)
	if err != nil {
		log.Error("Error -DB: User", err, row)
	}

	return "success"
}

//Get Home details based on Address
func Getfeedback_db() (Tempfeedarray []Model.Feedback) {

	var Data Model.Feedback
	rows, err := OpenConnection["Rentmatics"].Query("Select name,Emailid,rating from  feedback")
	if err != nil {
		log.Error("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Name,
			&Data.Emailid,
			&Data.Rating,
		)
		Tempfeedarray = append(Tempfeedarray, Data)
	}

	return Tempfeedarray
}

func InsertReferandearn(InsRef Model.Referandearn) string {
	row, err := OpenConnection["Rentmatics"].Exec("insert into referandearn (Referalname,Referalnumber,ownername,ownerphone,owneraddress) values (?,?,?,?,?)", InsRef.Refername, InsRef.Refernumber, InsRef.Ownername, InsRef.Ownerphone, InsRef.Owneraddress)
	if err != nil {
		log.Error("Error -DB: User", err, row)
	}

	return "success"
}

func Getcode_DB(InsRef string) string {
	var refercode string
	rows, err := OpenConnection["Rentmatics"].Query("select Referid from userdata where loginid=?", InsRef)
	if err != nil {
		log.Error("Error -DB: User", err, rows)
	}

	for rows.Next() {

		rows.Scan(
			&refercode,
		)

	}
	return refercode
}
