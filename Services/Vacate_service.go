package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func Getvacatedetails(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllVacatedetails()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Vacatesingle struct {
	vacateid int
}

func Getsinglevacate(w http.ResponseWriter, r *http.Request) {

	var GetVacate Vacatesingle
	err := json.NewDecoder(r.Body).Decode(&GetVacate)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetSinglevacate_Db(GetVacate.vacateid)

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
