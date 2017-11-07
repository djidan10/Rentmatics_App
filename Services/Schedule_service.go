package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ScheduleVisit(w http.ResponseWriter, r *http.Request) {

	var Schedule Model.Schedule

	err := json.NewDecoder(r.Body).Decode(&Schedule)
	if err != nil {
		log.Error("Error:Schedule Visit", err)
	}
	var Tostring = Schedule.Scheduleemail

	from := mail.NewEmail("Rentmatics User", "sandhiyabalakrishnan6@gmail.com")
	subject := "RENTMATICS NOTIFICATION - CONFORM SCHEDULE VISIT!"
	to := mail.NewEmail("Example User", Tostring)
	plainTextContent := "Your Appointment is Conformed,Our Executive will Contact you Soon\nThank you For Contacting Rentmatics"
	htmlContent := "<strong>Your Appointment is Conformed,Our Executive will Contact you Soon\nThank you For Contacting Rentmatics</strong><strong>Thank you for Contacting Rentmatics</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Error(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)

	}

	Db.Scheduleinsert(Schedule)

}
