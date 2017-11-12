package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPendingCR() (Temprentarray Model.Admincomp_Req) {

	Temprentarray.PendingComplaints = GetAllPendingComplaint()
	Temprentarray.PendingRequest = GetAllpendingrequest()
	return Temprentarray
}
func GetOwnerAllPendingCR(Ownerid int) (Temprentarray Model.Admincomp_Req) {

	Temprentarray.PendingComplaints = GetAllOwnerPendingComplaint(Ownerid)
	Temprentarray.PendingRequest = GetOwnerAllpendingrequest(Ownerid)
	return Temprentarray
}
