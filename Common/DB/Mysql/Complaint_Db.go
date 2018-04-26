package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"

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
			&Data.Raisedby,
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
			&Data.Raisedby,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

//Get All pending Complaint
func GetAllOwnerPendingComplaint(OTid int) (Temprentarray []Model.Complaints) {

	var Data Model.Complaints

	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where complaint_status = 'pending' and Raisedby = 'Owner' and tenantid =?", OTid)
	if err != nil {
		log.Error("Error -Db:Tenant Pending Complaint", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.Complaint_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Complaint_raisedDate,
			&Data.Raisedby,
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
			&Data.Raisedby,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)

	}
	Datasend.ComplaintsDetails = Data
	Datasend.Tenantdetails = GetIndivualTenant_Db(Data.Tenant_Id)
	return
}

//Get Simple Complaint
func GetSingleComplainttenant_Db(tenantid int) (Datasend Model.Complaintsend1) {
	var Data1 []Model.Complaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from complaint where Raisedby = 'Tenant' and tenantid =?", tenantid)
	if err != nil {
		log.Error("Error -Db:Tenant SingleComplaint", err)
	}

	for rows.Next() {
		fmt.Println("data")
		var Data Model.Complaints

		rows.Scan(

			&Data.Complaint_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Complaint_raisedDate,
			&Data.Raisedby,
			&Data.Complaint_Description,
			&Data.Complaint_status,
			&Data.Complaint_solvedDate,
		)
		Data1 = append(Data1, Data)

	}
	Datasend.ComplaintsDetails = Data1
	Datasend.Tenantdetails = GetIndivualTenant_Db(tenantid)
	return
}

func InsertComplaints_Db(Complaintdata Model.Complaints) {
	fmt.Println(Complaintdata.Raisedby, Complaintdata.Complaint_raisedDate, Complaintdata.Complaint_solvedDate, Complaintdata.Complaint_status)
	rows, err := OpenConnection["Rentmatics"].Exec("insert into complaint (homeid,tenantid,complained_raiseddate,Raisedby,complaint_description,complaint_status,complaint_solveddate) values (?,?,?,?,?,?,?)", Complaintdata.Home_id, Complaintdata.Tenant_Id, Complaintdata.Complaint_raisedDate, Complaintdata.Raisedby, Complaintdata.Complaint_Description, Complaintdata.Complaint_status, Complaintdata.Complaint_solvedDate)
	log.Info("successfully inserted", rows, err)

}

func UpdateComplaints_Db(Owncompupdatedata Model.Complaintstatus) {
	Queryupdate := "UPDATE complaint SET complaint_status=' " + Owncompupdatedata.Status + " ' , complaint_solveddate = '" + Owncompupdatedata.ApproveDate + "'  where id= '" + Owncompupdatedata.Complaintid + "'"

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	log.Info(rows, err)
}
