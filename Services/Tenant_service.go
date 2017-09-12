package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTenant(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllTenantdetails()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
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
		fmt.Println("err", err)
	}

	Data := Db.GetSingleTenant_Db(Gettenantid.Tenantid)

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
