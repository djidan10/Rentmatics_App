package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func InsertFav(Fav Model.Favourites) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into favourities (favid,homeid,userid) values (?,?,?)", Fav.Favid, Fav.Homeid, Fav.UserId)
	fmt.Println(rows, err)
}

//Get Home details based on Address
func GetFavourites(Favuser int) (Temprentarray []Model.RentSendFav) {

	var Data Model.Home
	var TempRentStruct Model.RentSendFav

	row, err := OpenConnection["Rentmatics"].Query("Select homeid from favourities where userid=?", Favuser)
	for row.Next() {

		var Favhome int
		row.Scan(
			&Favhome,
		)

		// query
		rows := OpenConnection["Rentmatics"].QueryRow("Select * from home where id=?", Favhome)
		fmt.Println(err)

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

		//		row := OpenConnection["Rentmatics"].QueryRow("Select cities from cities where cityid=?", Data.Cityid)

		//		var Cit string
		//		row.Scan(
		//			&Cit,
		//		)
		row1, err := OpenConnection["Rentmatics"].Query("select * from pictures_url where home_id=?", Data.Id)
		fmt.Println("err", err)
		var Rentimgarray []Model.Home_images

		for row1.Next() {
			var Rentimage Model.Home_images

			row1.Scan(

				&Rentimage.Home_id,
				&Rentimage.Picture_id,
				&Rentimage.Picture_url)

			Rentimgarray = append(Rentimgarray, Rentimage)
		}
		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray
		//TempRentStruct.RentCity = Cit

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}

//////Get Home details based on Address
//func GetFavourites(Favuser Model.GetFavourites.UserId) (Temprentarray []Model.RentSendFav) {

//	var Data Model.Home
//	var TempRentStruct Model.RentSendFav
//	var Favhome Model.GetFavourites.Homeid

//		// query
//		rows, err := OpenConnection["Rentmatics"].Query("Select homeid from favourities where userid=?", Favuser.UserId)
//		fmt.Println(err)

//		for rows.Next() {

//			rows.Scan(
//			&Favhome.Homeid,)

//			rows, err s= OpenConnection["Rentmatics"].Query("Select * from home where cityid=?", Favhome.Homeid)
//			fmt.Println(err)

//	for rows.Next() {

//		rows.Scan(
//			&Data.Id,
//			&Data.Housename,
//			&Data.Adress1,
//			&Data.Adress2,
//			&Data.Cityid,
//			&Data.State,
//			&Data.Country,
//			&Data.Pin,
//			&Data.Phone_number,
//			&Data.Month_rent,
//			&Data.Tenant_type,
//			&Data.Booking_type,
//			&Data.House_type,
//			&Data.Bhk,
//			&Data.Distance,
//			&Data.Furnish_type,
//			&Data.Secutity_deposit,
//		)
//		rows, err := OpenConnection["Rentmatics"].Query("select * from pictures_url where home_id=?", Data.Id)
//		fmt.Println("err", err)
//		var Rentimgarray []Model.Home_images

//		for rows.Next() {
//			var Rentimage Model.Home_images

//			rows.Scan(

//				&Rentimage.Home_id,
//				&Rentimage.Picture_id,
//				&Rentimage.Picture_url)

//			Rentimgarray = append(Rentimgarray, Rentimage)
//		}
//		TempRentStruct.RentFullStruct = Data
//		TempRentStruct.RentFullimages = Rentimgarray

//		Temprentarray = append(Temprentarray, TempRentStruct)
//	}
//}
//	return Temprentarray
//}
