package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"encoding/json"
	"net/http"
)

//Get All owner Complaint Detials
func GetOwnercomplaint(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllOwnercomplaint()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error on Get all Owner complaints", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Ownercompl struct {
	Ownercomid int
}

//Get Single OwnerDetails
func GetSingleOwnercomplaint(w http.ResponseWriter, r *http.Request) {

	var GetOwnerid Ownercompl
	err := json.NewDecoder(r.Body).Decode(&GetOwnerid)
	if err != nil {
		log.Error("Error :Single owner complaint", err)
	}
	Data := Db.GetSingleOwnercomplaint_Db(GetOwnerid.Ownercomid)
	Senddata, err1 := json.Marshal(Data)
	if err1 != nil {
		log.Error("Error :Single owner complaint from Database", err1)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
