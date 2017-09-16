package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func GetHomeDetilswithFav(Cityid string, Login string) []Model.RentSend {
	var homeidd []int
	var ResultFav []Model.RentSend

	rows, err := OpenConnection["Rentmatics"].Query("Select homeid from wishlist where loginid=?", Login)
	fmt.Println(err)

	for rows.Next() {
		var homeid int
		rows.Scan(
			&homeid,
		)

		homeidd = append(homeidd, homeid)
		fmt.Println("***********homeidd", homeidd)

	}

	ResultFav = GetHomeDetils_DB(Cityid)

	for k, v := range ResultFav {

		for _, v1 := range homeidd {
			if v.RentFullStruct.Id == v1 {

				//v.RentFullStruct.Liked = true
				ResultFav[k].RentFullStruct.Liked = true

			}

			fmt.Println("inside if ", v1, v.RentFullStruct.Id)
		}

	}
	return ResultFav
}

//Get Home details based on Address
func GetHomeDetils_DB(City string) (Temprentarray []Model.RentSend) {

	var Data Model.Home
	var TempRentStruct Model.RentSend

	//	City, _ := strconv.Atoi(Cit)
	//fmt.Println("id..............", City)
	rows, err := OpenConnection["Rentmatics"].Query("Select * from home where city=?", City)
	fmt.Println(err)

	for rows.Next() {

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

		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", City)
		//		row.Scan(&TempRentStruct.Cityname)

		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}

//Get single home based on id
func GetSinglehome_Db(homeid int) (Data1 Model.Home_single) {
	var Data Model.Home

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from home where id=?", homeid)
	fmt.Println(err)

	for rows.Next() {

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
		)
		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", Data.Cityid)
		//		row.Scan(&Data1.Cityname)
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

		Data1.Home_Data = Data
		Data1.Home_images = Rentimgarray
		fmt.Println("end of Data", Data1)

	}
	return
}

//Get single home based on id
func GetSinglehome_DbFav(homeid int, Login string) (Data1 Model.Home_single) {
	var Data Model.Home

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from home where id=?", homeid)
	fmt.Println(err)

	for rows.Next() {

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
		)
		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", Data.Cityid)
		//		row.Scan(&Data1.Cityname)
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

		var Count int
		query := "SELECT COUNT(*) FROM wishlist WHERE wishlist.loginid ='" + Login + "'" + "and wishlist.homeid = " + fmt.Sprintf("%v", Data.Id)
		row := OpenConnection["Rentmatics"].QueryRow(query)
		row.Scan(
			&Count,
		)
		if Count == 0 {
			Data1.Liked = false
		} else {
			Data1.Liked = true
		}

		Data1.Home_Data = Data
		Data1.Home_images = Rentimgarray
		fmt.Println("end of Data", Data1)

	}
	return
}

//Get detils based on filter option

func GetFilterwithFav(Filtstr Model.Filter) []Model.RentSend {
	var homeidd []int
	var ResultFav []Model.RentSend

	rows, err := OpenConnection["Rentmatics"].Query("Select homeid from wishlist where loginid=?", Filtstr.F_Loginid)
	fmt.Println(err)

	for rows.Next() {
		var homeid int
		rows.Scan(
			&homeid,
		)

		homeidd = append(homeidd, homeid)
		fmt.Println("***********homeidd", homeidd)

	}

	ResultFav = GetFilter_Db(Filtstr)

	for k, v := range ResultFav {

		for _, v1 := range homeidd {
			if v.RentFullStruct.Id == v1 {

				//v.RentFullStruct.Liked = true
				ResultFav[k].RentFullStruct.Liked = true

			}

			fmt.Println("inside if ", v1, v.RentFullStruct.Id)
		}

	}
	return ResultFav
}

func GetFilter_Db(Filt Model.Filter) (Temprentarray []Model.RentSend) {
	var Data Model.Home

	var TempRentStruct Model.RentSend

	fmt.Sprintf("%.6f", Filt.F_Monthrent_Min)
	Query := "Select * from home where month_rent BETWEEN " + fmt.Sprintf("%v", Filt.F_Monthrent_Min) + " and " + fmt.Sprintf("%v", Filt.F_Monthrent_Max) + " OR tenant_type='" + Filt.F_Tenant_type + "'  OR booking_type= '" + Filt.F_Bookingtype + "'  OR house_type= '" + Filt.F_Housetype + "'  OR furnish_type= '" + Filt.F_Furnished_type + "'  OR distance= '" + Filt.F_Distance + "'  OR bhk= " + fmt.Sprintf("%v", Filt.F_Bhk) + "  OR bed= " + fmt.Sprintf("%v", Filt.F_Bed)

	//	Query := "Select * from home where month_rent BETWEEN " + fmt.Sprintf("%v", Filt.F_Monthrent_Min) + " and " + fmt.Sprintf("%v", Filt.F_Monthrent_Max) + " and tenant_type='" + Filt.F_Tenant_type + "' and booking_type= '" + Filt.F_Bookingtype + "' and house_type= '" + Filt.F_Housetype + "' and furnish_type= '" + Filt.F_Furnished_type + "' and distance= '" + Filt.F_Distance + "' and bhk= " + fmt.Sprintf("%v", Filt.F_Bhk) + " and bed= " + fmt.Sprintf("%v", Filt.F_Bed)

	fmt.Println("querrr", Query)
	rows, err := OpenConnection["Rentmatics"].Query(Query)

	fmt.Println(err)

	for rows.Next() {

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
		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", Data.Cityid)
		//		row.Scan(&TempRentStruct.Cityname)

		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}

func GetallhomedetailsDB() (Temprentarray []Model.RentSend) {
	var Data Model.Home
	var TempRentStruct Model.RentSend

	rows, err := OpenConnection["Rentmatics"].Query("Select * from home")
	fmt.Println(err)

	for rows.Next() {

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

		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", Data.Cityid)
		//		row.Scan(&TempRentStruct.Cityname)

		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray

}

//Get Home details based on Address
func GetHomeDetils_DBwithOwnerid(Ownerid int) (Temprentarray []Model.RentSend) {

	var Data Model.Home
	var TempRentStruct Model.RentSend

	//	City, _ := strconv.Atoi(Cit)
	//fmt.Println("id..............", City)
	rows, err := OpenConnection["Rentmatics"].Query("Select * from home where ownerid=?", Ownerid)
	fmt.Println(err)

	for rows.Next() {

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

		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", City)
		//		row.Scan(&TempRentStruct.Cityname)

		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}

//Get Home details based on Address
func GetHomeDetils_DBwithExecutiveid(Executiveid int) (Temprentarray []Model.RentSendTenant) {

	var Data Model.Home
	var TempRentStruct Model.RentSendTenant

	//	City, _ := strconv.Atoi(Cit)
	//fmt.Println("id..............", City)
	rows, err := OpenConnection["Rentmatics"].Query("Select * from home where executiveid=?", Executiveid)
	fmt.Println(err)

	for rows.Next() {

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

		//		row := OpenConnection["Rentmatics"].QueryRow("select cities from cities where id=?", City)
		//		row.Scan(&TempRentStruct.Cityname)

		TempRentStruct.RentFullStruct = Data
		TempRentStruct.RentFullimages = Rentimgarray
		TempRentStruct.Tenantdetails = GetIndivualTenant_Db(Data.Executive_id)

		Temprentarray = append(Temprentarray, TempRentStruct)
	}

	return Temprentarray
}
