package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllExecutive() (Temprentarray []Model.Executivedetails) {

	var Data Model.Executivedetails
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  executive")
	if err != nil {
		log.Error("Error -DB: All Executibe", err)
	}

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
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  executive where id=?", executiveid)
	if err != nil {
		log.Error("Error -DB: All Executibe", err)
	}
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

func GetIndivualExecutive_Db(executiveid int) (Data Model.Executivedetails) {

	rows, err := OpenConnection["Rentmatics"].Query("Select * from  executive where id=?", executiveid)
	if err != nil {
		log.Error("Error -DB: All Executibe", err)
	}

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

	return
}

func Inserthome_DB(Homedata Model.HomeInsert, Imageurl []string) {

	fmt.Println("inside execitive")
	Executiveid, _ := strconv.Atoi(Homedata.ExecutiveID)
	Ownerid, _ := strconv.Atoi(Homedata.Ownerid)
	Bed, _ := strconv.Atoi(Homedata.Bed)
	Bhk, _ := strconv.Atoi(Homedata.Bhk)
	Bookbed, _ := strconv.Atoi(Homedata.Booked_bed)
	Availbed, _ := strconv.Atoi(Homedata.Availble_Bed)
	Availroom, _ := strconv.Atoi(Homedata.Avail_room)
	Bookbhk, _ := strconv.Atoi(Homedata.Booked_bhk)
	Monthlyrent, _ := strconv.ParseFloat(Homedata.Month_rent, 32)
	Bhkrent, _ := strconv.ParseFloat(Homedata.Bhk_rent, 32)
	Bedrent, _ := strconv.ParseFloat(Homedata.Bed_rent, 32)
	Securitydeposit, _ := strconv.ParseFloat(Homedata.Secutity_deposit, 32)

	rows, err := OpenConnection["Rentmatics"].Exec("insert into home (housename,executiveid,ownerid,adress1,adress2,city,district,state,country,pin,phonenumber,month_rent,bed_rent,bhk_rent,tenant_type,booking_type,house_type,bhk,bed,Avail_bed,Avail_room,Booked_bed,Booked_bhk,distance,furnish_type,secutity_deposit,listing,Amenities,Description,latitude,longitude,Squarefeet,Totalfloors,Facing,Parking) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", Homedata.Housename, Executiveid, Ownerid, Homedata.Adress1, Homedata.Adress2, Homedata.City, Homedata.District, Homedata.State, Homedata.Country, Homedata.Pin, Homedata.Phone_number, Monthlyrent, Bedrent, Bhkrent, Homedata.Tenant_type, Homedata.Booking_type, Homedata.House_type, Bhk, Bed, Availbed, Availroom, Bookbed, Bookbhk, Homedata.Distance, Homedata.Furnish_type, Securitydeposit, Homedata.Listing, Homedata.Amenities, Homedata.Description, Homedata.Latitude, Homedata.Longitude, Homedata.Squarefeet, Homedata.Totalfloors, Homedata.Facing, Homedata.Parking)

	if err != nil {
		log.Error("Error -DB: Executive insert", err)
	}

	var pictureurl Model.Home_images
	pictureurl.Home_id, _ = rows.LastInsertId()

	for _, pictureurl.Picture_url = range Imageurl {
		rows, err := OpenConnection["Rentmatics"].Exec("insert into pictures_url (home_id,image) values (?,?)", pictureurl.Home_id, pictureurl.Picture_url)
		if err != nil {
			log.Error("Error -DB: Executive insert picture", err, rows)
		}
	}

}

//Get Home details based on Address
func Checklogin_DB(User Model.Login) (Uservalid Model.LoginData) {

	rows, err := OpenConnection["Rentmatics"].Query("Select * from roles where username=?", User.Username)
	if err != nil {
		log.Error("Error -DB: Check login", err)
	}

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

func InsertExecutive_Db(Executivedata Model.Executivedetails) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into executive (First_Name,Last_Name,Email_Id,Contact,Alernate_Contact,DOB,Permanent_Address1,Permanent_Address2,Permanent_Area,Permanent_City,Permanent_Pin,Pan_Card,Aadhar_Card,Voter_Card) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", Executivedata.First_Name, Executivedata.Last_Name, Executivedata.Email_Id, Executivedata.Contact, Executivedata.Alernate_Contact, Executivedata.DOB, Executivedata.Permanent_Address1, Executivedata.Permanent_Address2, Executivedata.Permanent_Area, Executivedata.Permanent_City, Executivedata.Permanent_Pin, Executivedata.Pan_Card, Executivedata.Aadhar_Card, Executivedata.Voter_Card)
	log.Info("successfully inserted", rows, err)

}
