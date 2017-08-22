package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"

	"encoding/json"
	"fmt"
	"net/http"
)

func GetCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside city")

	Data := Db.GetAllCities()
	//fmt.Println("data is ", data)

	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
