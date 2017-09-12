package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllOwner() (Temprentarray []Model.Owner) {

	var Data Model.Owner
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  owner")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Owner_id,
			&Data.Homeid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Permanent_Address1,
			&Data.Permanent_Address2,
			&Data.Permanent_Area,
			&Data.Permanent_City,
			&Data.Permanent_Pin,
			&Data.Tenant_img,
			&Data.Pan_Card,
			&Data.Aadhar_Card,
			&Data.Voter_Card,
			&Data.Aggrement,
		)

		Temprentarray = append(Temprentarray, Data)
	}

	return Temprentarray
}

func GetSingleOwner_Db(Ownerid int) (Data Model.Owner) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  owner where id=?", Ownerid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Owner_id,
			&Data.Homeid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Permanent_Address1,
			&Data.Permanent_Address2,
			&Data.Permanent_Area,
			&Data.Permanent_City,
			&Data.Permanent_Pin,
			&Data.Tenant_img,
			&Data.Pan_Card,
			&Data.Aadhar_Card,
			&Data.Voter_Card,
			&Data.Aggrement,
		)

	}
	return
}
