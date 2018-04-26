package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetComplaint(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllComplaint()

	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func Getpendingstatus(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllPendingComplaint()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
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
		log.Error("Error :tenant Complaint", err)
	}

	Data := Db.GetSingleComplaint_Db(GetComplaintid.Complaintid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func Getsingletenantcomplaint(w http.ResponseWriter, r *http.Request) {

	var GetComplaintid Complaintviewid
	err := json.NewDecoder(r.Body).Decode(&GetComplaintid)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}

	Data := Db.GetSingleComplainttenant_Db(GetComplaintid.Complaintid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error :tenant Complaint", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func InsertComplaints(w http.ResponseWriter, r *http.Request) {

	var InsComplaint Model.Complaints
	err := json.NewDecoder(r.Body).Decode(&InsComplaint)
	if err != nil {
		log.Error("Error: Tenant vacate ", err)
	}
	//Insert Vacate Details
	fmt.Println(InsComplaint.Complaint_status)
	Db.InsertComplaints_Db(InsComplaint)

}
