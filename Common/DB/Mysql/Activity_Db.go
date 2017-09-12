package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllActivitydetails() (Temprentarray []Model.Activity) {

	var Data Model.Activity
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  activities")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Activity_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Activity_Date,
			&Data.Activity_Tittle,
			&Data.Participation_count,
			&Data.Activity_Description,
			&Data.Activity_Status,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleActivity_Db(Activityid int) (Data Model.Activity) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  activities where id=?", Activityid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Activity_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Activity_Date,
			&Data.Activity_Tittle,
			&Data.Participation_count,
			&Data.Activity_Description,
			&Data.Activity_Status,
		)

	}
	return
}
