package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

func InsertWishlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside Rent inssert")
	var Wish Model.Wishlist
	err := json.NewDecoder(r.Body).Decode(&Wish)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(Wish)

	Db.Insertwishlist(Wish)

}

type Favourite struct {
	Loginid string
}

func DeleteWishlist(w http.ResponseWriter, r *http.Request) {

	var WishDel Model.WishlistDelete
	err := json.NewDecoder(r.Body).Decode(&WishDel)
	if err != nil {
		fmt.Println("err", err)
	}

	Db.Deletewishlist(WishDel)

}

func Wishlist(w http.ResponseWriter, r *http.Request) {
	var Fav Favourite
	fmt.Println("inside Wishlist")

	err := json.NewDecoder(r.Body).Decode(&Fav)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(",,,,,,,,,,,,,,,", Fav.Loginid)

	Data := Db.GetWishDB(Fav.Loginid)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}
