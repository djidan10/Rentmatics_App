package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"

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
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Raisedby,
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
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Raisedby,
			&Data.Approve_Date,
			&Data.Status,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

//Get Home details based on Address
func GetOwnerAllpendingrequest(Ownerid int) (Temprentarray []Model.Request) {

	var Data Model.Request
	rows, err := OpenConnection["Rentmatics"].Query("Select * from request where status = 'pending' and Raisedby = 'Owner' and tenantid =?", Ownerid)

	if err != nil {
		log.Error("Error -DB: Tenant Request", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.Request_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Resquestername,
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Raisedby,
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
			&Data.Request_Date,
			&Data.Request_description,
			&Data.Raisedby,
			&Data.Approve_Date,
			&Data.Status,
		)

	}

	Datasend.Resquestdetails = Data
	Datasend.Tenantdetails = GetIndivualTenant_Db(Data.Tenant_Id)
	return
}

func InsertRequest_Db(Requestdata Model.Request) {
	fmt.Println(Requestdata.Request_description)
	rows, err := OpenConnection["Rentmatics"].Exec("insert into request (homeid,tenantid,Resquestername,Request_Date,Request_description,Raisedby,Approve_Date,Status) values (?,?,?,?,?,?,?,?)", Requestdata.Home_id, Requestdata.Tenant_Id, Requestdata.Resquestername, Requestdata.Request_Date, Requestdata.Request_description, Requestdata.Raisedby, Requestdata.Approve_Date, Requestdata.Status)
	log.Info("successfully inserted", rows, err)

}

func UpdateRequest_Db(requpdatedata Model.Requeststatus) {
	Queryupdate := "UPDATE request SET Status=' " + requpdatedata.Status + " ' , Approve_Date = '" + requpdatedata.ApproveDate + "'  where id= '" + requpdatedata.Requestid + "'"

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	log.Info(rows, err)
}
