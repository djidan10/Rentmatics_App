package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

//Get All owners Details
func GetOwners(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllOwner()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error:Get All Owner", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Ownersingle struct {
	Ownerid int
}

//Get Single Owner Details
func GetsingleOwner(w http.ResponseWriter, r *http.Request) {

	var GetOwner Ownersingle
	err := json.NewDecoder(r.Body).Decode(&GetOwner)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Data := Db.GetSingleOwner_Db(GetOwner.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

//Get Single Owner Details
func Ownertoexecutive(w http.ResponseWriter, r *http.Request) {

	var GetOwnerExecutive Model.OwnertoExecutive
	err := json.NewDecoder(r.Body).Decode(&GetOwnerExecutive)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Db.GetOwnerExecutive_Db(GetOwnerExecutive)

}

//Get Single Owner Details
func Insertowner(w http.ResponseWriter, r *http.Request) {

	var InsOwner Model.Owner
	err := json.NewDecoder(r.Body).Decode(&InsOwner)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Db.Insertowner_Db(InsOwner)

}
