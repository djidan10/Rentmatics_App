package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
)

func Notification(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Rent")

	var Notify Model.Notification
	err := json.NewDecoder(r.Body).Decode(&Notify)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("********", Notify)
	var ToString = Notify.NotifyEmail

	auth := smtp.PlainAuth("", "Rentmatics@gmail.com", "RENTMATICS2017", "smtp.gmail.com")
	var receiptent = "To:" + ToString + "\r\n"
	to := []string{ToString}
	msg := []byte(receiptent +
		"Subject: RENTMATICS NOTIFICATION!\r\n" +
		"\r\n" +
		"Currently We are working under this area,Rentmatics comre to you as soon!\r\n")

	mailerr := smtp.SendMail("smtp.gmail.com:587", auth, "Rentmatics@gmail.com", to, msg)

	if err != nil {
		fmt.Println(mailerr)
	}
	fmt.Println("finished mailing")

	Db.Insertnotify_DB(Notify)

}
