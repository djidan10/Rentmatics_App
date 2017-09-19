package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

type Cityy struct {
	Id      string
	Loginid string
}

//Get All Home Detials without service id
func Getallhomedetails(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetallhomedetailsDB()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Cannot fetch Data All home details", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

//Gethome detaiils based on city
func Gethomedetails(w http.ResponseWriter, r *http.Request) {
	var Cityid Cityy
	var Data []Model.RentSend
	err := json.NewDecoder(r.Body).Decode(&Cityid)
	if err != nil {
		log.Error("Error:Marshal Data", err)
	}

	if Cityid.Loginid == "" {
		Data = Db.GetHomeDetils_DB(Cityid.Id)

	} else {
		Data = Db.GetHomeDetilswithFav(Cityid.Id, Cityid.Loginid)

	}
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error on send home details", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type HomeInd struct {
	Id      int
	Loginid string
}

//get Indivual home
func GetIndivual_home(w http.ResponseWriter, r *http.Request) {
	var Data Model.Home_single
	var Homeid HomeInd
	err := json.NewDecoder(r.Body).Decode(&Homeid)
	if err != nil {
		log.Error("Error-Indivual home", err)
	}
	if Homeid.Loginid == "" {
		Data = Db.GetSinglehome_Db(Homeid.Id)
	} else {
		Data = Db.GetSinglehome_DbFav(Homeid.Id, Homeid.Loginid)
	}
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error-Indivual home", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func GetFilter(w http.ResponseWriter, r *http.Request) {
	var Filt Model.Filter
	err := json.NewDecoder(r.Body).Decode(&Filt)
	if err != nil {
		log.Error("Error on Filterhome", err)
	}
	Data := Db.GetFilterwithFav(Filt)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error on Filterhome", err)

	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}
