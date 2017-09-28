package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

func Userdata(w http.ResponseWriter, r *http.Request) {
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
	cookie := &http.Cookie{
		Name:  "RentmaticsCookie",
		Value: Data.Loginid,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

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
	if err != nil {
		log.Error("Error - User Login", err)
	}
	if Data.Status == "Success" {

		cookie := &http.Cookie{
			Name:  "RentmaticsCookie",
			Value: Data.Loginid,
			Path:  "/",
		}
		var senddata Model.StatusResponse

		http.SetCookie(w, cookie)
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
