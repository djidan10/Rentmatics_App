package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"

	"encoding/json"

	"net/http"
)

func GetComplaintRequest(w http.ResponseWriter, r *http.Request) {
	//Connecting Database
	Data := Db.GetAllPendingCR()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Cannot unmarshal the Activity data from database", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type OwnerComplaintrequest struct {
	Ownerid int
}

func GetOwnerComplaintRequest(w http.ResponseWriter, r *http.Request) {
	var GetRC OwnerComplaintrequest
	err := json.NewDecoder(r.Body).Decode(&GetRC)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Connecting Database
	Data := Db.GetOwnerAllPendingCR(GetRC.Ownerid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Cannot unmarshal the Activity data from database", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
