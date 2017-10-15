package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Get Home details based on Address
func GetAllVacatedetails() (Temprentarray []Model.Vacate) {

	var Data Model.Vacate
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  vacate")
	if err != nil {
		log.Error("Error -DB: Vacate", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.Vacate_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Vacate_priordate,
			&Data.Vacate_Date,
			&Data.Vacate_reason,
			&Data.Vacate_status,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSinglevacate_Db(vacateid int) (Data Model.Vacate) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from vacate where id=?", vacateid)
	if err != nil {
		log.Error("Error -DB: Vacate", err)
	}

	for rows.Next() {

		rows.Scan(
			&Data.Vacate_id,
			&Data.Home_id,
			&Data.Tenant_Id,
			&Data.Vacate_priordate,
			&Data.Vacate_Date,
			&Data.Vacate_reason,
			&Data.Vacate_status,
		)

	}
	return
}

func Insertvacate_Db(Vacatedata Model.Vacate) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into vacate (homeid,tenantid,Vacate_predate,Vacate_Date,Vacate_reason,Vacate_status) values (?,?,?,?,?,?)", Vacatedata.Home_id, Vacatedata.Tenant_Id, Vacatedata.Vacate_priordate, Vacatedata.Vacate_Date, Vacatedata.Vacate_reason, Vacatedata.Vacate_status)
	log.Info("successfully inserted", rows, err)

}
