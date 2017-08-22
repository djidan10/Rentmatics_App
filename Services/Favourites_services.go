package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func Getfavourites(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Rent")

	var Fav Model.GetFavourites
	err := json.NewDecoder(r.Body).Decode(&Fav)
	if err != nil {
		fmt.Println("err", err)
	}

	Data := Db.GetFavourites(Fav.UserId)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)
}

func InsertFavourite(w http.ResponseWriter, r *http.Request) {
	var Favinsert Model.Favourites
	err := json.NewDecoder(r.Body).Decode(&Favinsert)
	if err != nil {
		fmt.Println("err", err)
	}

	Db.InsertFav(Favinsert)
}
