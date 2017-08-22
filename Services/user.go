package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func Userdata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Rent")
	var User1 Model.User
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("*********", User1)

	Data := Db.InserUser(User1)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
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
		fmt.Println("err", err)
	}
	fmt.Println("*********", User)

	Data := Db.GetUSer(User)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}
	if Data.Status == "Success" {

		fmt.Println("inside cookie", Data.Loginid)
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

	var User1 Model.User
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("*********", User1)

	Data := Db.InserUser(User1)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
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
