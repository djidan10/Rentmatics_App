package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"

	"fmt"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Get All owners Details
func GetOwners(w http.ResponseWriter, r *http.Request) {
	Data := Db.GetAllOwner()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error:Get All Owner", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Ownersingle struct {
	Ownerid int
}

//Get Single Owner Details
func GetsingleOwner(w http.ResponseWriter, r *http.Request) {

	var GetOwner Ownersingle
	err := json.NewDecoder(r.Body).Decode(&GetOwner)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Data := Db.GetSingleOwner_Db(GetOwner.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

//Get Single Owner Details
func Ownertoexecutive(w http.ResponseWriter, r *http.Request) {

	var GetOwnerExecutive Model.OwnertoExecutive
	err := json.NewDecoder(r.Body).Decode(&GetOwnerExecutive)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Tostring := GetOwnerExecutive.Email
	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("Rentmatics User", "Rentmatics@gmail.com")
	subject := "RENTMATICS NOTIFICATION!"
	to := mail.NewEmail("Example User", Tostring)
	plainTextContent := "We Have Succesfully Recieved your information,Our Rentmatics Executive will Contact you Soon \nThank you For Contacting Rentmatics"
	htmlContent := "<strong>RENTMATICS NOTIFICATION! </strong><br><br><p>We Have Succeesfully Recieved your home information,Our Executive will Contact you Soon</p><br><br><strong>Thank you For Contacting Rentmatics</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(sendkey)
	response, err := client.Send(message)
	if err != nil {
		log.Error(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)

	}
	log.Info("successfully send Email")
	Db.GetOwnerExecutive_Db(GetOwnerExecutive)

}

//Get Single Owner Details
func Insertowner(w http.ResponseWriter, r *http.Request) {

	var InsOwner Model.Owner
	err := json.NewDecoder(r.Body).Decode(&InsOwner)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Db.Insertowner_Db(InsOwner)

}

type Ownerhome struct {
	Ownerid int
}

//Get Single Owner Details
func Getmyhomedetails(w http.ResponseWriter, r *http.Request) {

	var Ownerhome1 Ownerhome
	err := json.NewDecoder(r.Body).Decode(&Ownerhome1)
	if err != nil {
		log.Error("Error - Owner single", err)
	}

	Data := Db.GetHomeDetils_DBwithOwnerid(Ownerhome1.Ownerid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

//Get Single Owner Details
func InsertOwnerRequestDetails(w http.ResponseWriter, r *http.Request) {

	var InsOwner1 Model.OwnerRequest
	err := json.NewDecoder(r.Body).Decode(&InsOwner1)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Db.InsertownerRequest_Db(InsOwner1)

}

//Get Single Owner Details
func InsertOwnerVisitDetails(w http.ResponseWriter, r *http.Request) {

	var InsOwner2 Model.OwnerVisit
	err := json.NewDecoder(r.Body).Decode(&InsOwner2)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Db.InsertownerVisitRequest_Db(InsOwner2)

}

//Get Single Owner Details
func InsertOwnerMessage(w http.ResponseWriter, r *http.Request) {

	var InsOwner3 Model.Ownermessage
	err := json.NewDecoder(r.Body).Decode(&InsOwner3)
	if err != nil {
		log.Error("Error - Owner single", err)
	}

	Db.InsertownerMessage_Db(InsOwner3)

}

type Getownerid struct {
	Ownerdata string
}

//Get Single Owner Details
func GetsingleId(w http.ResponseWriter, r *http.Request) {

	var GetOwnerid Getownerid
	err := json.NewDecoder(r.Body).Decode(&GetOwnerid)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Data := Db.GetOwnerId_Db(GetOwnerid.Ownerdata)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type GetownerMessage struct {
	Email string
}

//Get Single Owner Details
func GetOwnermessage(w http.ResponseWriter, r *http.Request) {

	var GetOwneremailRole GetownerMessage
	err := json.NewDecoder(r.Body).Decode(&GetOwneremailRole)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Data := Db.GetEmailRole_Db(GetOwneremailRole.Email)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

type Messagedata struct {
	Messageid int
}

//Get Single Owner Details
func GetOwnermessagedetails(w http.ResponseWriter, r *http.Request) {

	var GetOwneremailRole Messagedata
	err := json.NewDecoder(r.Body).Decode(&GetOwneremailRole)
	if err != nil {
		log.Error("Error - Owner single", err)
	}
	Data := Db.GetMessageId_Db(GetOwneremailRole.Messageid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - Get single Owner Data from DB", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
