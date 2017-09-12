package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetActivity(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllActivitydetails()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Activitysingle struct {
	Activityid int
}

func GetsingleActivity(w http.ResponseWriter, r *http.Request) {

	var GetActivity Activitysingle
	err := json.NewDecoder(r.Body).Decode(&GetActivity)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetSingleActivity_Db(GetActivity.Activityid)

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
