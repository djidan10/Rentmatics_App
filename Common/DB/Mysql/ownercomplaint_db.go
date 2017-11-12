package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllOwnercomplaint() (Temprentarray []Model.Ownercomplaints) {

	var Data Model.Ownercomplaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint")
	if err != nil {
		log.Error("Error -DB: All Owners", err)
	}

	for rows.Next() {

		rows.Scan(

			&Data.Ownercom_id,
			&Data.Ownerid,
			&Data.Homeid,
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

	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint where Id=?", Ownerid)
	if err != nil {
		log.Error("Error -DB: Get Single owner", err)
	}
	for rows.Next() {

		rows.Scan(

			&Data.Ownercom_id,
			&Data.Ownerid,
			&Data.Homeid,
			&Data.ComplaintDate,
			&Data.Message,
			&Data.Status,
			&Data.Approvedate,
		)

	}

	return
}

//Get Home details based on Address
func GetOwnerComplaint(Complaintsowner int) (Temprentarray []Model.Ownercomplaints) {

	var Data Model.Ownercomplaints
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint where Ownerid=?", Complaintsowner)
	if err != nil {
		log.Error("Error -DB: All Owners", err)
	}

	for rows.Next() {

		rows.Scan(

			&Data.Ownercom_id,
			&Data.Ownerid,
			&Data.Homeid,
			&Data.ComplaintDate,
			&Data.Message,
			&Data.Status,
			&Data.Approvedate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

//Get Home details based on Address
func GetOwnerRequest(Complaintsowner int) (Temprentarray []Model.OwnerRequest) {

	var Data Model.OwnerRequest
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownerrequest where ownerid=?", Complaintsowner)
	if err != nil {
		log.Error("Error -DB: All Owners", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.OwnerRequestid,
			&Data.OwnerId,
			&Data.Ownername,
			&Data.Owneremail,
			&Data.Ownerhomeid,
			&Data.Description,
			&Data.Status,
			&Data.RequestDate,
			&Data.ApproveDate,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

//Get Home details based on Address
func GetSingleOwnerRequest(Complaintsowner int) (Data Model.OwnerRequest) {

	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownerrequest where ownerrequestid=?", Complaintsowner)
	if err != nil {
		log.Error("Error -DB: All Owners", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.OwnerRequestid,
			&Data.OwnerId,
			&Data.Ownername,
			&Data.Owneremail,
			&Data.Ownerhomeid,
			&Data.Description,
			&Data.Status,
			&Data.RequestDate,
			&Data.ApproveDate,
		)

	}

	return Data
}
func Insertownercomplaints_Db(Owncompdata Model.Ownercomplaints) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into ownercomplaint (Ownerid,Homeid,ComplaintDate,Message,Status,Approvedate) values (?,?,?,?,?,?)", Owncompdata.Ownerid, Owncompdata.Homeid, Owncompdata.ComplaintDate, Owncompdata.Message, Owncompdata.Status, Owncompdata.Approvedate)
	log.Info(rows, err)
}
