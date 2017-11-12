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

type TenantOwnerhome struct {
	Ownerid int
}

func GetOwnerTenantDetails(w http.ResponseWriter, r *http.Request) {

	var Ownertenant TenantOwnerhome
	err := json.NewDecoder(r.Body).Decode(&Ownertenant)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetOwnerTenantdetails(Ownertenant.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func GetOwnercomplaintId(w http.ResponseWriter, r *http.Request) {

	var Ownergetid TenantOwnerhome
	err := json.NewDecoder(r.Body).Decode(&Ownergetid)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetOwnerComplaint(Ownergetid.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func GetOwnerResquestId(w http.ResponseWriter, r *http.Request) {

	var Ownergetid TenantOwnerhome
	err := json.NewDecoder(r.Body).Decode(&Ownergetid)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetOwnerRequest(Ownergetid.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func GetSingleOwnerResquest(w http.ResponseWriter, r *http.Request) {

	var Ownergetid TenantOwnerhome
	err := json.NewDecoder(r.Body).Decode(&Ownergetid)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetSingleOwnerRequest(Ownergetid.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Tenantstring struct {
	Tenantid string
}

func GetAdminTenant(w http.ResponseWriter, r *http.Request) {

	var Gettenantidd Tenantstring
	err := json.NewDecoder(r.Body).Decode(&Gettenantidd)
	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	Data := Db.GetAdminTenant_Db(Gettenantidd.Tenantid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :Get Tenant", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
