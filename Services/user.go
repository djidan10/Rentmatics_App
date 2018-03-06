package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func Userdata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("insert user")
	var User1 Model.User
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	Data := Db.InserUser(User1)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - User Login", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func RentUserdata(w http.ResponseWriter, r *http.Request) {
	var Rentuser Model.RentUser
	err := json.NewDecoder(r.Body).Decode(&Rentuser)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	Data := Db.InserRentUser(Rentuser)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - User Login", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func RentUserlogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "RentmaticsAdminCookie",
		Value: "NODATA",
		Path:  "/",
	}
	http.SetCookie(w, cookie)

}
func Userlogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "RentmaticsCookie",
		Value: "NODATA",
		Path:  "/",
	}
	http.SetCookie(w, cookie)

}
func Login(w http.ResponseWriter, r *http.Request) {
	var User Model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - User Login", err)
	}

	Data := Db.GetUSer(User)
	Senddata, err := json.Marshal(Data)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func RentLogin(w http.ResponseWriter, r *http.Request) {

	var User Model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	fmt.Println("rentlogin", User.Password, User.Username)
	Data := Db.GetRentUSer(User)
	fmt.Println(Data)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	if Data.Role != "Invalid" {

		var senddata Model.StatusResponse
		senddata.Status = "Success"
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")

		w.Write(Senddata)

	} else {

		w.Write(Senddata)

	}
}

func Changepassword(w http.ResponseWriter, r *http.Request) {

	var User1 Model.Changepass
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		log.Error("Error - Change password", err)
	}
	Data := Db.InsertChangepassword(User1)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Forgotpassword struct {
	Email string
}

func Forgot(w http.ResponseWriter, r *http.Request) {

	var Emailget Forgotpassword
	err := json.NewDecoder(r.Body).Decode(&Emailget)
	if err != nil {
		log.Error("Error - Change password", err)
	}
	Data := Db.InsertForgotspassword(Emailget.Email)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func Feedback(w http.ResponseWriter, r *http.Request) {

	var Getfeedback Model.Feedback
	err := json.NewDecoder(r.Body).Decode(&Getfeedback)
	if err != nil {
		log.Error("Error - Feedback", err)
	}
	Data := Db.Insertfeedback(Getfeedback)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func GetFeedback(w http.ResponseWriter, r *http.Request) {

	Data := Db.Getfeedback_db()

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func Referandearn(w http.ResponseWriter, r *http.Request) {

	var GetRefer Model.Referandearn
	err := json.NewDecoder(r.Body).Decode(&GetRefer)
	if err != nil {
		log.Error("Error - Feedback", err)
	}
	Data := Db.InsertReferandearn(GetRefer)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func GetRefercode(w http.ResponseWriter, r *http.Request) {

	var GetRefer Model.Refer
	err := json.NewDecoder(r.Body).Decode(&GetRefer)
	if err != nil {
		log.Error("Error - Feedback", err)
	}
	Data := Db.Getcode_DB(GetRefer.Loginid)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
