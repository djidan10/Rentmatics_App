package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetCity_DB(Userdata Model.Login) (Uservalid Model.Login) {

	// query
	//var Uservalid Model.Login
	rows, err := OpenConnection["Rentmatics"].Query("Select userdata from username where username=?", Userdata.Username)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Uservalid.Username,
			//&Uservalid.password,
		)

	}

	return Uservalid
}

func GetAllCities() []Model.CityStruct {
	var TempCitySlice []Model.CityStruct
	rows, err := OpenConnection["Rentmatics"].Query("Select * from cities")
	fmt.Println(err)
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
