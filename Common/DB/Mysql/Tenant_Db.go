package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

//Get Home details based on Address
func GetAllTenantdetails() (Temprentarray []Model.Tenant) {

	var Data Model.Tenant
	rows, err := OpenConnection["Rentmatics"].Query("Select * from tenant")
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Tenant_id,
			&Data.Homeid,
			&Data.Executiveid,
			&Data.Ownerid,
			&Data.Loginid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Tenant_Type,
			&Data.Office_Address1,
			&Data.Office_Address2,
			&Data.Office_Area,
			&Data.Office_City,
			&Data.Office_Pin,
			&Data.Office_Email,
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

func GetSingleTenant_Db(Tenantid int) (Datasend Model.Tenantsend) {
	var Data Model.Tenant
	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from tenant where id=?", Tenantid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Tenant_id,
			&Data.Homeid,
			&Data.Executiveid,
			&Data.Ownerid,
			&Data.Loginid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Tenant_Type,
			&Data.Office_Address1,
			&Data.Office_Address2,
			&Data.Office_Area,
			&Data.Office_City,
			&Data.Office_Pin,
			&Data.Office_Email,
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

	row1, err := OpenConnection["Rentmatics"].Query("select * from tenant_payment where tenant_id=?", Data.Tenant_id)
	fmt.Println("err", err)
	var Tenantpay1 []Model.Tenantpayment

	for row1.Next() {
		var Tenatpay Model.Tenantpayment

		row1.Scan(

			&Tenatpay.Tenantpaymentid,
			&Tenatpay.Tenant_id,
			&Tenatpay.Home_id,
			&Tenatpay.Owner_id,
			&Tenatpay.Loginid,
			&Tenatpay.Executiveid,
			&Tenatpay.Lastmonth_paiddate,
			&Tenatpay.Duedate,
			&Tenatpay.RentAmount,
			&Tenatpay.Status,
			&Tenatpay.PaymentDetails,
			&Tenatpay.TransactionDetails,
		)

		Tenantpay1 = append(Tenantpay1, Tenatpay)
	}

	Data1 := GetSinglehome_Db(Data.Homeid)
	Datasend.Tenantdetails = Data
	Datasend.Homedetails = Data1
	Datasend.TenantPaymentdetails = Tenantpay1
	fmt.Println(Data1)

	return
}

func GetIndivualTenant_Db(Tenantid int) (Data Model.Tenant) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from tenant where id=?", Tenantid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Data.Tenant_id,
			&Data.Homeid,
			&Data.Executiveid,
			&Data.Ownerid,
			&Data.Loginid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Email_Id,
			&Data.Contact,
			&Data.Alernate_Contact,
			&Data.DOB,
			&Data.Tenant_Type,
			&Data.Office_Address1,
			&Data.Office_Address2,
			&Data.Office_Area,
			&Data.Office_City,
			&Data.Office_Pin,
			&Data.Office_Email,
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
