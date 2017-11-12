package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllRequestdetails()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :Get Request", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func GetPendingrequest(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllpendingrequest()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Requestsingle struct {
	Requestid int
}

func GetsingleRequest(w http.ResponseWriter, r *http.Request) {

	var GetRequestsingle Requestsingle
	err := json.NewDecoder(r.Body).Decode(&GetRequestsingle)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}

	Data := Db.GetSingleRequest_Db(GetRequestsingle.Requestid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func InsertRequest(w http.ResponseWriter, r *http.Request) {

	var InsRequest Model.Request
	err := json.NewDecoder(r.Body).Decode(&InsRequest)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.InsertRequest_Db(InsRequest)

}
func Updaterequeststatus(w http.ResponseWriter, r *http.Request) {

	var UpRequest Model.Requeststatus
	err := json.NewDecoder(r.Body).Decode(&UpRequest)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	Db.UpdateRequest_Db(UpRequest)

}
