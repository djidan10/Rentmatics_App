package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

func InsertWishlist(w http.ResponseWriter, r *http.Request) {

	var Wish Model.Wishlist
	err := json.NewDecoder(r.Body).Decode(&Wish)
	if err != nil {
		log.Error("Error : Wishlist ", err)
	}
	Db.Insertwishlist(Wish)

}

type Favourite struct {
	Loginid string
}

func DeleteWishlist(w http.ResponseWriter, r *http.Request) {

	var WishDel Model.WishlistDelete
	err := json.NewDecoder(r.Body).Decode(&WishDel)
	if err != nil {
		log.Error("Error: Delete  Wishlist ", err)
	}
	Db.Deletewishlist(WishDel)

}

func Wishlist(w http.ResponseWriter, r *http.Request) {
	var Fav Favourite
	err := json.NewDecoder(r.Body).Decode(&Fav)
	if err != nil {
		log.Error("Error: Wishlist", err)
	}
	Data := Db.GetWishDB(Fav.Loginid)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error: Wishlist ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
