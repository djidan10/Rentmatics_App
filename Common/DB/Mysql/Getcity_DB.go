package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetCity_DB(Userdata Model.Login) (Uservalid Model.Login) {

	rows, err := OpenConnection["Rentmatics"].Query("Select userdata from username where username=?", Userdata.Username)
	if err != nil {
		log.Error("Error -DB: Get City", err)
	}

	for rows.Next() {

		rows.Scan(
			&Uservalid.Username,
		)

	}

	return Uservalid
}

func GetAllCities() []Model.CityStruct {
	var TempCitySlice []Model.CityStruct
	rows, err := OpenConnection["Rentmatics"].Query("Select * from cities")
	if err != nil {
		log.Error("Error -DB: Get All cities", err)
	}
	for rows.Next() {
		var TempCityStruct Model.CityStruct
		rows.Scan(
			&TempCityStruct.Id,
			&TempCityStruct.City,
		)
		TempCitySlice = append(TempCitySlice, TempCityStruct)
	}
	return TempCitySlice
}
