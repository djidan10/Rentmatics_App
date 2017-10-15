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

func InsertRequest_Db(Requestdata Model.Request) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into request (homeid,tenantid,Resquestername,Total_Request,Pending_Request,Solved_Request,Request_Date,Request_description,Approve_Date,Status) values (?,?,?,?,?,?,?,?,?,?)", Requestdata.Home_id, Requestdata.Tenant_Id, Requestdata.Resquestername, Requestdata.Total_Request, Requestdata.Pending_Request, Requestdata.Solved_Request, Requestdata.Request_Date, Requestdata.Request_description, Requestdata.Approve_Date, Requestdata.Status)
	log.Info("successfully inserted", rows, err)

}
