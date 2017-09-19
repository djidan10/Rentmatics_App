package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
	"net/smtp"
)

//Notify When Home Available
func Notification(w http.ResponseWriter, r *http.Request) {
	var Notify Model.Notification
	err := json.NewDecoder(r.Body).Decode(&Notify)
	if err != nil {
		log.Error("Notification Error", err)
	}
	var ToString = Notify.NotifyEmail
	auth := smtp.PlainAuth("", "Rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
	var receiptent = "To:" + ToString + "\r\n"
	to := []string{ToString}
	msg := []byte(receiptent +
		"Subject: RENTMATICS NOTIFICATION!\r\n" +
		"\r\n" +
		"Currently We are working under this area,Rentmatics comre to you as soon!\r\n")

	mailerr := smtp.SendMail("smtp.gmail.com:587", auth, "Rentmatics@gmail.com", to, msg)
	if mailerr != nil {
		log.Error("Mail Error - Notification", mailerr)

	}
	log.Info("successfully send Email")
	Db.Insertnotify_DB(Notify)

}
