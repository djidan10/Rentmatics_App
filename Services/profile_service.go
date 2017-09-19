package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

//Profile Insert - User
func Profileinsert(w http.ResponseWriter, r *http.Request) {

	var Profile1 Model.Profile
	err := json.NewDecoder(r.Body).Decode(&Profile1)
	if err != nil {
		log.Error("Error on profileinsert", err)
	}
	Data := Db.Insertprofile(Profile1)
	Senddata, err := json.Marshal(Data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

//GetProfile - User
func Getprofile(w http.ResponseWriter, r *http.Request) {

	var GetProfiledata Model.Getprofile
	err := json.NewDecoder(r.Body).Decode(&GetProfiledata)
	if err != nil {
		log.Error("Struct Mismatch on get profile", err)
	}
	Data := Db.Getprofile(GetProfiledata)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Getprofile error from database", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
