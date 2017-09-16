package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllComplaint() (Temprentarray []Model.Complaints) {

	var Data Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Complaint_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Complaint_raisedDate,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetAllPendingComplaint() (Temprentarray []Model.Complaints) {

	var Data Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where complaint_status =?", "pending")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Complaint_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Complaint_raisedDate,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}
func GetSingleComplaint_Db(Complaintid int) (Data Model.Complaints) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where id=?", Complaintid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Complaint_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Complaint_raisedDate,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)

	}
	return
}
