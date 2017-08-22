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

func Getallhomedetails(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetallhomedetailsDB()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func Gethomedetails(w http.ResponseWriter, r *http.Request) {

	var Cityid Cityy
	var Data []Model.RentSend
	err := json.NewDecoder(r.Body).Decode(&Cityid)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(Cityid)
	if Cityid.Loginid == "" {
		fmt.Println("inside no fav")
		Data = Db.GetHomeDetils_DB(Cityid.Id)

	} else {
		fmt.Println("inside fave")
		Data = Db.GetHomeDetilswithFav(Cityid.Id, Cityid.Loginid)

	}
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type HomeInd struct {
	Id      int
	Loginid string
}

func GetIndivual_home(w http.ResponseWriter, r *http.Request) {
	var Data Model.Home_single
	var Homeid HomeInd
	err := json.NewDecoder(r.Body).Decode(&Homeid)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(Homeid.Id)
	if Homeid.Loginid == "" {
		Data = Db.GetSinglehome_Db(Homeid.Id)
	} else {
		Data = Db.GetSinglehome_DbFav(Homeid.Id, Homeid.Loginid)

	}
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
func GetFilter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***************************************")

	var Filt Model.Filter
	err := json.NewDecoder(r.Body).Decode(&Filt)
	if err != nil {
		fmt.Println("err", err)
	}
	Data := Db.GetFilterwithFav(Filt)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)
}
