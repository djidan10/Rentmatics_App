package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetComplaint(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllComplaint()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
func Getpendingstatus(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllPendingComplaint()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Complaintviewid struct {
	Complaintid int
}

func GetsingleComplaint(w http.ResponseWriter, r *http.Request) {

	var GetComplaintid Complaintviewid
	err := json.NewDecoder(r.Body).Decode(&GetComplaintid)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetSingleComplaint_Db(GetComplaintid.Complaintid)

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
