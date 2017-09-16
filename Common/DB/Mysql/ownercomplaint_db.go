package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllOwnercomplaint() (Temprentarray []Model.Ownercomplaints) {

	var Data Model.Ownercomplaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Ownercom_id,
			&Data.Ownerid,
			&Data.Executiveid,
			&Data.Tenantid,
			&Data.Homeid,
			&Data.Loginid,
			&Data.ComplaintDate,
			&Data.Message,
			&Data.Status,
			&Data.Approvedate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleOwnercomplaint_Db(Ownerid int) (Data Model.Ownercomplaints) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint where Id=?", Ownerid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Ownercom_id,
			&Data.Ownerid,
			&Data.Executiveid,
			&Data.Tenantid,
			&Data.Homeid,
			&Data.Loginid,
			&Data.ComplaintDate,
			&Data.Message,
			&Data.Status,
			&Data.Approvedate,
		)

	}

	return
}
