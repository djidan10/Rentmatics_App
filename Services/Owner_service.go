package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOwners(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllOwner()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Ownersingle struct {
	Ownerid int
}

func GetsingleOwner(w http.ResponseWriter, r *http.Request) {

	var GetOwner Ownersingle
	err := json.NewDecoder(r.Body).Decode(&GetOwner)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetSingleOwner_Db(GetOwner.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
