package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"encoding/json"
	"net/http"
)

func GetTenant(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllTenantdetails()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Tenantsingle struct {
	Tenantid int
}

func GetSingleTenant(w http.ResponseWriter, r *http.Request) {

	var Gettenantid Tenantsingle
	err := json.NewDecoder(r.Body).Decode(&Gettenantid)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetSingleTenant_Db(Gettenantid.Tenantid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
