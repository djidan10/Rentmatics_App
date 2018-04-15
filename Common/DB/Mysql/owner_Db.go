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
			&Data.Rusername,
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
			&Data.Rusername,
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
func GetOwnerId_Db(Ownername string) (Data Model.Owner) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select * from  owner where RMusername=?", Ownername)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.Owner_id,
			&Data.Homeid,
			&Data.First_Name,
			&Data.Last_Name,
			&Data.Rusername,
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

func GetEmailRole_Db(Email string) (Temprentarray []Model.ResponseMessage) {
	var Datavalue Model.ResponseMessage

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select id,subject,homeid,message  from  ownermessage where toemail=?", Email)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(
			&Datavalue.MessageId,
			&Datavalue.Subject,
			&Datavalue.Homeid,
			&Datavalue.Message,
		)
		Temprentarray = append(Temprentarray, Datavalue)

	}
	return Temprentarray
}

func GetMessageId_Db(Messageid int) (Data Model.ResponseMessage) {

	// query
	rows, err := OpenConnection["Rentmatics"].Query("Select id,subject,homeid,message  from  ownermessage where id=?", Messageid)
	fmt.Println(err)

	for rows.Next() {

		rows.Scan(

			&Data.MessageId,
			&Data.Subject,
			&Data.Homeid,
			&Data.Message,
		)

	}
	return
}

//func Inserthome_DB(Homedata Model.HomeInsert, Imageurl []string) {
func GetOwnerExecutive_Db(Homedata Model.OwnertoExecutive) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into ownertoexecutive (Name,Email,phone,Tenanttype,Address) values (?,?,?,?,?)", Homedata.Name, Homedata.Email, Homedata.Phone, Homedata.Tenanttype, Homedata.Address)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}

//func Inserthome_DB(Homedata Model.HomeInsert, Imageurl []string) {
func Getcontact_Db(Homedata Model.Contactform) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into contact (name,mail,number,subject,message) values (?,?,?,?,?)", Homedata.Name, Homedata.Email, Homedata.Phonenumber, Homedata.Subject, Homedata.Message)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}

//func Inserthome_DB(Homedata Model.HomeInsert, Imageurl []string) {
func Insertowner_Db(Insertownerdata Model.Owner) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into owner (Homeid,First_Name,Last_Name,Email_Id,Contact,Alernate_Contact,DOB,Permanent_Address1,Permanent_Address2,Permanent_Area,Permanent_City,Permanent_Pin,Tenant_img,Pan_Card,Aadhar_Card,Voter_Card,Aggrement) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", Insertownerdata.Homeid, Insertownerdata.First_Name, Insertownerdata.Last_Name, Insertownerdata.Email_Id, Insertownerdata.Contact, Insertownerdata.Alernate_Contact, Insertownerdata.DOB, Insertownerdata.Permanent_Address1, Insertownerdata.Permanent_Address2, Insertownerdata.Permanent_Area, Insertownerdata.Permanent_City, Insertownerdata.Permanent_Pin, Insertownerdata.Tenant_img, Insertownerdata.Pan_Card, Insertownerdata.Aadhar_Card, Insertownerdata.Voter_Card, Insertownerdata.Aggrement)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}

func InsertownerRequest_Db(Insertownerdata Model.OwnerRequest) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into ownerrequest (ownerid,ownername,owneremail,Homeid,Description,status) values (?,?,?,?,?,?)", Insertownerdata.OwnerId, Insertownerdata.Ownername, Insertownerdata.Owneremail, Insertownerdata.Ownerhomeid, Insertownerdata.Description, Insertownerdata.Status)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}

func InsertownerVisitRequest_Db(Insertownerdata Model.OwnerVisit) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into  ownervisit (ownername,owneremail,homeid,visitdate,description,status) values (?,?,?,?,?,?)", Insertownerdata.Ownername, Insertownerdata.Owneremail, Insertownerdata.Ownerhomeid, Insertownerdata.VisitDate, Insertownerdata.Description, Insertownerdata.Status)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}

func InsertownerMessage_Db(Insertownermasg Model.Ownermessage) {
	fmt.Println(Insertownermasg.Homeid, Insertownermasg.Description, Insertownermasg.Fromemail, Insertownermasg.Toemail)

	rows, err := OpenConnection["Rentmatics"].Exec("insert into  ownermessage (fromemail,toemail,subject,homeid,Towhom,message) values (?,?,?,?,?,?)", Insertownermasg.Fromemail, Insertownermasg.Toemail, Insertownermasg.Subject, Insertownermasg.Homeid, Insertownermasg.Towhom, Insertownermasg.Description)
	if err != nil {
		log.Error("Error -DB: Executive insert", rows)
	}

}
