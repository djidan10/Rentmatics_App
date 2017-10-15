package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllComplaint() (Temprentarray []Model.Complaints) {

	var Data Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint")
	if err != nil {
		log.Error("Error -Db:Tenant Complaint", err)
	}
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

//Get All pending Complaint
func GetAllPendingComplaint() (Temprentarray []Model.Complaints) {

	var Data Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where complaint_status =?", "pending")
	if err != nil {
		log.Error("Error -Db:Tenant Pending Complaint", err)
	}

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

//Get Simple Complaint
func GetSingleComplaint_Db(Complaintid int) (Datasend Model.Complaintsend) {
	var Data Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where id=?", Complaintid)
	if err != nil {
		log.Error("Error -Db:Tenant SingleComplaint", err)
	}

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
	Datasend.ComplaintsDetails = Data
	Datasend.Tenantdetails = GetIndivualTenant_Db(Data.Tenant_Id)
	return
}

func InsertComplaints_Db(Complaintdata Model.Complaints) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into complaint (homeid,tenantid,complained_raiseddate,complaint_description,complaint_status,complaint_solveddate) values (?,?,?,?,?,?)", Complaintdata.Home_id, Complaintdata.Tenant_Id, Complaintdata.Complaint_raisedDate, Complaintdata.Complaint_Description, Complaintdata.Complaint_status, Complaintdata.Complaint_solvedDate)
	log.Info("successfully inserted", rows, err)

}
