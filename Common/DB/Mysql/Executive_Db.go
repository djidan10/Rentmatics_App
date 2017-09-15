package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	"strconv"

	//"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllExecutive() (Temprentarray []Model.Executivedetails) {

	var Data Model.Executivedetails
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  executive")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Executive_id,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Permanent_Address1,
			&Data.Permanent_Address2,
			&Data.Permanent_Area,
			&Data.Permanent_City,
			&Data.Permanent_Pin,
			&Data.Executive_img,
			&Data.Pan_Card,
			&Data.Aadhar_Card,
			&Data.Voter_Card,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleExecutive_Db(executiveid int) (Datasend Model.Executivesend) {

	var Data Model.Executivedetails
	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  executive where id=?", executiveid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Executive_id,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Permanent_Address1,
			&Data.Permanent_Address2,
			&Data.Permanent_Area,
			&Data.Permanent_City,
			&Data.Permanent_Pin,
			&Data.Executive_img,
			&Data.Pan_Card,
			&Data.Aadhar_Card,
			&Data.Voter_Card,
		)

	}
	GetAllhome := GetHomeDetils_DBwithExecutiveid(Data.Executive_id)
	Datasend.ExecutiveData = Data
	Datasend.HomeData = GetAllhome
	return
}

func Inserthome_DB(Homeinsert Model.HomeInsert, Imageurl []string) {
	var Homedata Model.Home
	fmt.Println("inside  insidekfhlsalhf", Homeinsert.City)

	// Get City ID
	//row := OpenConnection["Rentmatics"].QueryRow("Select  from cities where cities=?", Homeinsert.City)

	//	var Cit int
	//	row.Scan(
	//		&Cit,
	//	)
	Homedata.Housename = Homeinsert.Housename
	Homedata.Adress1 = Homeinsert.Adress1
	Homedata.Adress2 = Homeinsert.Adress2
	Homedata.City = Homeinsert.City
	Homedata.State = Homeinsert.State
	Homedata.Country = Homeinsert.Country
	Homedata.Pin, _ = strconv.Atoi(Homeinsert.Pin)
	Homedata.Phone_number = Homeinsert.Phone_number
	Homedata.Month_rent, _ = strconv.ParseFloat(Homeinsert.Month_rent, 5)
	Homedata.Bed_rent, _ = strconv.ParseFloat(Homeinsert.Bed_rent, 5)
	Homedata.Month_rent, _ = strconv.ParseFloat(Homeinsert.Bhk_rent, 5)
	Homedata.Tenant_type = Homeinsert.Tenant_type
	Homedata.Booking_type = Homeinsert.Booking_type
	Homedata.House_type = Homeinsert.House_type
	Homedata.Bhk, _ = strconv.Atoi(Homeinsert.Bhk)
	Homedata.Bed, _ = strconv.Atoi(Homeinsert.Bed)
	Homedata.Furnish_type = Homeinsert.Furnish_type
	Homedata.Secutity_deposit, _ = strconv.ParseFloat(Homeinsert.Secutity_deposit, 5)
	Homedata.Distance = Homeinsert.Distance
	Homedata.Listing = Homeinsert.Listing
	Homedata.Amenities = Homeinsert.Amenities
	Homedata.Description = Homeinsert.Description

	rows, err := OpenConnection["Rentmatics"].Exec("insert into renthome (housename,adress1,adress2,city,state,country,pin,phonenumber,month_rent,tenant_type,booking_type,house_type,bhk,bed,distance,furnish_type,secutity_deposit,listing,Amenities,Description) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", Homedata.Housename, Homedata.Adress1, Homedata.Adress2, Homedata.City, Homedata.State, Homedata.Country, Homedata.Pin, Homedata.Phone_number, Homedata.Month_rent, Homedata.Tenant_type, Homedata.Booking_type, Homedata.House_type, Homedata.Bhk, Homedata.Bed, Homedata.Distance, Homedata.Furnish_type, Homedata.Secutity_deposit, Homedata.Listing, Homedata.Amenities, Homedata.Description)
	fmt.Println("successfully inserted", rows, err)

	var pictureurl Model.Home_images
	pictureurl.Home_id, _ = rows.LastInsertId()

	for _, pictureurl.Picture_url = range Imageurl {
		rows, err := OpenConnection["Rentmatics"].Exec("insert into pictures_url (home_id,image) values (?,?)", pictureurl.Home_id, pictureurl.Picture_url)
		fmt.Println("successfully inserted", rows, err)
	}

}

//Get Home details based on Address
func Checklogin_DB(User Model.Login) (Uservalid Model.LoginData) {

	rows, err := OpenConnection["Rentmatics"].Query("Select * from roles where username=?", User.Username)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Uservalid.Id,
			&Uservalid.Username,
			&Uservalid.Password,
			&Uservalid.Role,
		)

	}

	return Uservalid

}
