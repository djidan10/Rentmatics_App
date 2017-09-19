package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
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
