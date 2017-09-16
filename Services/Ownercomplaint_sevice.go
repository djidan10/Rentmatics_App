package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOwnercomplaint(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllOwnercomplaint()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Ownercompl struct {
	Ownercomid int
}

func GetSingleOwnercomplaint(w http.ResponseWriter, r *http.Request) {

	var GetOwnerid Ownercompl
	err := json.NewDecoder(r.Body).Decode(&GetOwnerid)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetSingleOwnercomplaint_Db(GetOwnerid.Ownercomid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
