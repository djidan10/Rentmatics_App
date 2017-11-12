package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Cityy struct {
	Id      string
	Loginid string
}

type LoginHome struct {
	Loginid string
}

//Get All Home Detials without service id
func Getallhomedetails(w http.ResponseWriter, r *http.Request) {

	var LoginData LoginHome
	var Data []Model.RentSend
	err := json.NewDecoder(r.Body).Decode(&LoginData)
	if err != nil {
		log.Error("Error:Marshal Data", err)
	}

	if LoginData.Loginid == "" {
		Data = Db.GetallhomedetailsDB()

	} else {
		Data = Db.GetAllHomeDetilswithFav(LoginData.Loginid)

	}
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error on send home details", err)
	}
	w.WriteHeader(http.StatusOK)
	//	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//	w.Header().Set("Access-Control-Allow-Orgin", "localhost")

	fmt.Println("header", w.Header())
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
		fmt.Println("inside non fav")

		Data = Db.GetSinglehome_Db(Homeid.Id)
	} else {

		fmt.Println("inside  fav")
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
	log.Info("inside filter")

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
