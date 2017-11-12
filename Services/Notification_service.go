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

//Notify When Home Available
func Notification(w http.ResponseWriter, r *http.Request) {
	var Notify Model.Notification
	err := json.NewDecoder(r.Body).Decode(&Notify)
	if err != nil {
		log.Error("Notification Error", err)
	}

	Tostring := Notify.NotifyEmail
	sendkey := os.Getenv("SENDGRID_API_KEYGO")

	from := mail.NewEmail("Rentmatics User", "Rentmatics@gmail.com")
	subject := "RENTMATICS NOTIFICATION!"
	to := mail.NewEmail("Example User", Tostring)
	plainTextContent := "We Have Succesfully Recieved your information,Rentmatics will Contact you Soon \nThank you For Contacting Rentmatics"
	htmlContent := "<strong>We Have Succeesfully Recieved your home information,Our Executive will Contact you Soon</strong><strong>Thank you For Contacting Rentmatics</strong>"
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
	Db.Insertnotify_DB(Notify)

}
