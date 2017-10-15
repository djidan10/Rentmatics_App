package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"Rentmatics_App/Logger"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

var (
	log = Logger.NewLogger("Rentmaticsservice")
)

func GetActivity(w http.ResponseWriter, r *http.Request) {
	//Connecting Database
	Data := Db.GetAllActivitydetails()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Cannot unmarshal the Activity data from database", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Activitysingle struct {
	Activityid int
}

func GetsingleActivity(w http.ResponseWriter, r *http.Request) {

	var GetActivity Activitysingle
	err := json.NewDecoder(r.Body).Decode(&GetActivity)
	if err != nil {
		log.Error("Error on Get particular details", err)
	}
	//Connecting Database
	Data := Db.GetSingleActivity_Db(GetActivity.Activityid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func InsertActivity(w http.ResponseWriter, r *http.Request) {

	var InsActivity Model.Activity
	err := json.NewDecoder(r.Body).Decode(&InsActivity)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.InsertActivity_Db(InsActivity)

}
