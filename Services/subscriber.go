package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	//	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Subscribeuser struct {
	Sub_Email string
}

func Postsubscriber(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Subscriber")
	var User Subscribeuser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("*********", User)

	Db.Insertsubscribtion(User.Sub_Email)

}
