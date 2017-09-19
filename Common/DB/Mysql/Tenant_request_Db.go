package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllRequestdetails() (Temprentarray []Model.Request) {

	var Data Model.Request
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  request")
	if err != nil {
		log.Error("Error -DB: Tenant Request", err)
	}

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

//Get Home details based on Address
func GetAllpendingrequest() (Temprentarray []Model.Request) {

	var Data Model.Request
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  request where status=?", "pending")
	if err != nil {
		log.Error("Error -DB: Tenant Request", err)
	}

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

func GetSingleRequest_Db(Requestid int) (Datasend Model.Requestsend) {
	var Data Model.Request
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  request where id=?", Requestid)
	if err != nil {
		log.Error("Error -DB: Tenant Request", err)
	}

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
	Datasend.Resquestdetails = Data
	Datasend.Tenantdetails = GetIndivualTenant_Db(Data.Tenant_Id)
	return
}
