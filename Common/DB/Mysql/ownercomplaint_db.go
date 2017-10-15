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

	rows, err := OpenConnection["Rentmatics"].Query("Select * from  ownercomplaint where Id=?", Ownerid)
	if err != nil {
		log.Error("Error -DB: Get Single owner", err)
	}
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

func Insertownercomplaints_Db(Owncompdata Model.Ownercomplaints) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into ownercomplaint (Ownerid,Executiveid,Tenantid,Homeid,Loginid,ComplaintDate,Message,Status,Approvedate) values (?,?,?,?,?,?,?,?,?)", Owncompdata.Ownerid, Owncompdata.Executiveid, Owncompdata.Tenantid, Owncompdata.Homeid, Owncompdata.Loginid, Owncompdata.ComplaintDate, Owncompdata.Message, Owncompdata.Status, Owncompdata.Approvedate)
	log.Info(rows, err)
}
