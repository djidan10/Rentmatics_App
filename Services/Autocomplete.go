package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	"encoding/json"
	"net/http"
)

func GetCities(w http.ResponseWriter, r *http.Request) {

	//Calling Database
	Data := Db.GetAllCities()
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
