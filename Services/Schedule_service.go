package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
	"net/smtp"
)

func ScheduleVisit(w http.ResponseWriter, r *http.Request) {

	var Schedule Model.Schedule

	err := json.NewDecoder(r.Body).Decode(&Schedule)
	if err != nil {
		log.Error("Error:Schedule Visit", err)
	}
	var Tostring = Schedule.Scheduleemail
	auth := smtp.PlainAuth("", "Rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
	var receiptent = "To:" + Tostring + "\r\n"
	to := []string{Tostring}
	msg := []byte(receiptent +
		"Subject: RENTMATICS NOTIFICATION FOR YOUR APPOINTMENT!\r\n" +
		"\r\n" +
		"Your Appointment is Conformed,Our Executive will Contact you Soon!\r\n")

	mailerr := smtp.SendMail("smtp.gmail.com:587", auth, "Rentmatics@gmail.com", to, msg)
	if mailerr != nil {
		log.Error("Error:Sending schedule Email", mailerr)
	}
	log.Info("finished mailing")

	Db.Scheduleinsert(Schedule)

}
