package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"encoding/json"
	"net/http"
)

type Subscribeuser struct {
	Sub_Email string
}

func Postsubscriber(w http.ResponseWriter, r *http.Request) {

	var User Subscribeuser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - Subscrbibe")
	}

	Db.Insertsubscribtion(User.Sub_Email)

}
