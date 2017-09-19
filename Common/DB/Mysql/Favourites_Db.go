package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InsertFav(Fav Model.Favourites) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into favourities (favid,homeid,userid) values (?,?,?)", Fav.Favid, Fav.Homeid, Fav.UserId)
	if err != nil {
		log.Error("Error -DB: Insert Fav", err, rows)
	}
}

//Get Home details based on Address
func GetFavourites(Favuser int) (Temprentarray []Model.RentSendFav) {

	var Data Model.Home
	var TempRentStruct Model.RentSendFav

	row, err := OpenConnection["Rentmatics"].Query("Select homeid from favourities where userid=?", Favuser)
	if err != nil {
		log.Error("Error -DB: Get Favourites", err)
	}
	for row.Next() {

		var Favhome int
		row.Scan(
			&Favhome,
		)
		rows := OpenConnection["Rentmatics"].QueryRow("Select * from home where id=?", Favhome)

		rows.Scan(
			&Data.Id,
			&Data.Housename,
			&Data.Executive_id,
			&Data.Ownerid,
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
			&Data.Avail_bed,
			&Data.Avail_room,
			&Data.Booked_bed,
			&Data.Booked_bhk,
			&Data.Booked_room,
			&Data.Booked_home,
			&Data.Booked_Allbeds,
			&Data.Booked_AllRooms,
			&Data.Distance,
			&Data.Furnish_type,
			&Data.Secutity_deposit,
			&Data.Listing,
			&Data.Amenities,
			&Data.Description,
			&Data.Latitude,
			&Data.Longitude,
			&Data.Squarefeet,
			&Data.Likecount,
			&Data.Rating,
			&Data.Totalfloors,
			&Data.Facing,
			&Data.Parking,
		)

		row1, err := OpenConnection["Rentmatics"].Query("select * from pictures_url where home_id=?", Data.Id)
		if err != nil {
			log.Error("Error -DB: get Fav Picture URL", err)
		}
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

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}
