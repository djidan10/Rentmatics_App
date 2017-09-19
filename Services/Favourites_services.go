package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

//Send Favourited Details
func Getfavourites(w http.ResponseWriter, r *http.Request) {
	var Fav Model.GetFavourites
	err := json.NewDecoder(r.Body).Decode(&Fav)
	if err != nil {
		log.Error("Cannot Get value Favourite", err)
	}
	Data := Db.GetFavourites(Fav.UserId)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Info("Cannot unmarshal the data from Favourites", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

//Insert Favourite
func InsertFavourite(w http.ResponseWriter, r *http.Request) {
	var Favinsert Model.Favourites
	err := json.NewDecoder(r.Body).Decode(&Favinsert)
	if err != nil {
		fmt.Println("err", err)
	}
	Db.InsertFav(Favinsert)
}
