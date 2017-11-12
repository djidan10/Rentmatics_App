package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"Rentmatics_App/Logger"
)

var (
	log = Logger.NewLogger("RentmaticsDB")
)

//Get Home details based on Address
func GetAllActivitydetails() (Temprentarray []Model.Activity) {

	var Data Model.Activity
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  activities")
	if err != nil {
		log.Error("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Activity_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Activity_Date,
			&Data.Activity_Tittle,
			&Data.Participation_count,
			&Data.Activity_Description,
			&Data.Activity_Status,
			&Data.ApproveDate,
		)
		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}
func GetAllActivityPendingdetails() (Temprentarray []Model.Activity) {

	var Data Model.Activity
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  activities where Activity_Status=?", "pending")
	if err != nil {
		log.Error("Error -Db:Activity", err)
	}
	for rows.Next() {

		rows.Scan(
			&Data.Activity_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Activity_Date,
			&Data.Activity_Tittle,
			&Data.Participation_count,
			&Data.Activity_Description,
			&Data.Activity_Status,
			&Data.ApproveDate,
		)
		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleActivity_Db(Activityid int) (Data Model.Activity) {

	rows, err := OpenConnection["Rentmatics"].Query("Select * from  activities where id=?", Activityid)
	if err != nil {
		log.Error("Error -Db:Activity")
	}

	for rows.Next() {
		rows.Scan(
			&Data.Activity_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Activity_Date,
			&Data.Activity_Tittle,
			&Data.Participation_count,
			&Data.Activity_Description,
			&Data.Activity_Status,
			&Data.ApproveDate,
		)

	}
	return
}

func InsertActivity_Db(Activitydata Model.Activity) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into activities (homeid,tenantid,activity_date,Activity_Tittle,Participation_count,Activity_Description,Activity_Status,Approvedate) values (?,?,?,?,?,?,?,?,?)", Activitydata.Home_id, Activitydata.Tenant_Id, Activitydata.Activity_Date, Activitydata.Activity_Tittle, Activitydata.Participation_count, Activitydata.Activity_Description, Activitydata.Activity_Status, Activitydata.ApproveDate)
	log.Info("successfully inserted", rows, err)

}

func UpdateActivity_Db(Actupdatedata Model.ActivityUpdatestatus) {
	Queryupdate := "UPDATE activities SET Activity_Status=' " + Actupdatedata.Status + " ' , Approvedate = '" + Actupdatedata.ApproveDate + "'  where id= '" + Actupdatedata.Activityid + "'"

	rows, err := OpenConnection["Rentmatics"].Exec(Queryupdate)
	log.Info(rows, err)
}
