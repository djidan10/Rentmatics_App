package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	//	"strconv"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Insertwishlist(Wishinfo Model.Wishlist) {
	var Count int

	//query := "SELECT COUNT(*) FROM wishlist WHERE wishlist.loginid ='" + Wishinfo.Loginid + "' and wishlist.homeid = " + Wishinfo.Homeid
	query := "SELECT COUNT(*) FROM wishlist WHERE wishlist.loginid ='" + Wishinfo.Loginid + "' and wishlist.homeid = " + fmt.Sprintf("%v", Wishinfo.Homeid)

	row := OpenConnection["Rentmatics"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {
		row, err := OpenConnection["Rentmatics"].Exec("insert into wishlist (homeid,loginid) values (?,?)", Wishinfo.Homeid, Wishinfo.Loginid)
		fmt.Println("successfully inserted", row, err)

	}
}

//Get Home details based on Address
func GetWishDB(Loginid string) (Temprentarray []Model.RentSend) {

	var Homeid string

	var Data Model.Home
	Data.Liked = true
	var TempRentStruct Model.RentSend

	rows, err := OpenConnection["Rentmatics"].Query("Select homeid from wishlist where loginid=?", Loginid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Homeid,
		)

		rows, err1 := OpenConnection["Rentmatics"].Query("Select * from home where id=?", Homeid)
		fmt.Println(err1)

		for rows.Next() {

			rows.Scan(
				&Data.Id,
				&Data.Housename,
				&Data.Adress1,
				&Data.Adress2,
				&Data.City,
				&Data.District,
				&Data.State,
				&Data.Country,
				&Data.Pin,
				&Data.Phone_number,
				&Data.Month_rent,
				&Data.Bed_rent,
				&Data.Bhk_rent,
				&Data.Tenant_type,
				&Data.Booking_type,
				&Data.House_type,
				&Data.Bhk,
				&Data.Bed,
				&Data.Distance,
				&Data.Furnish_type,
				&Data.Secutity_deposit,
				&Data.Listing,
				&Data.Amenities,
				&Data.Description,
				&Data.Latitude,
				&Data.Longitude,
			)

			rows, err := OpenConnection["Rentmatics"].Query("select * from pictures_url where home_id=?", Data.Id)
			fmt.Println("err", err)
			var Rentimgarray []Model.Home_images

			for rows.Next() {
				var Rentimage Model.Home_images

				rows.Scan(

					&Rentimage.Home_id,
					&Rentimage.Picture_id,
					&Rentimage.Picture_url)

				Rentimgarray = append(Rentimgarray, Rentimage)
			}

			//			row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", Data.Cityid)
			//			row.Scan(&TempRentStruct.Cityname)

			TempRentStruct.RentFullStruct = Data
			TempRentStruct.RentFullimages = Rentimgarray

			Temprentarray = append(Temprentarray, TempRentStruct)
		}
	}
	return Temprentarray
}

func Deletewishlist(Wishinfo Model.WishlistDelete) {
	Query := "delete from wishlist where homeid=" + strconv.Itoa(Wishinfo.Homeid) + "  and loginid='" + Wishinfo.Loginid + "'"
	fmt.Println("inside delete", Query)

	row, err := OpenConnection["Rentmatics"].Exec(Query)
	fmt.Println("successfully Deleted", row, err)

}
