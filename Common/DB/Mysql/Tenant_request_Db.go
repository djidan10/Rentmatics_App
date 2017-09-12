package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllRequestdetails() (Temprentarray []Model.Request) {

	var Data Model.Request
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  request")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Request_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Resquestername,
			&Data.Total_Request,
			&Data.Pending_Request,
			&Data.Solved_Request,
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Approve_Date,
			&Data.Status,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleRequest_Db(Requestid int) (Data Model.Request) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  request where id=?", Requestid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Request_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Resquestername,
			&Data.Total_Request,
			&Data.Pending_Request,
			&Data.Solved_Request,
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Approve_Date,
			&Data.Status,
		)

	}
	return
}
