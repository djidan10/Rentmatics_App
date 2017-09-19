package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"encoding/json"
	"net/http"
)

func Getvacatedetails(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllVacatedetails()

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Vacatesingle struct {
	Vacateid int
}

func Getsinglevacate(w http.ResponseWriter, r *http.Request) {

	var GetVacate Vacatesingle
	err := json.NewDecoder(r.Body).Decode(&GetVacate)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	Data := Db.GetSinglevacate_Db(GetVacate.Vacateid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
