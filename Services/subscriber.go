package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"fmt"

	"encoding/json"
	"net/http"
)

func Postsubscriber(w http.ResponseWriter, r *http.Request) {

	var User Model.Subscribeuser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - Subscrbibe")
	}
	fmt.Println(User.Sub_City, User.Sub_Email)

	Db.Insertsubscribtion(User)

}
