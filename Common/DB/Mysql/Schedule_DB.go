package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func Scheduleinsert(Scheduledata Model.Schedule) {

	rows, err := OpenConnection["Rentmatics"].Exec("insert into schedule (homeid,name,phonenumber,email,date,time) values (?,?,?,?,?,?)", Scheduledata.Schedulehomeid, Scheduledata.ScheduleName, Scheduledata.Schedulephone, Scheduledata.Scheduleemail, Scheduledata.Scheduledate, Scheduledata.Scheduletime)
	fmt.Println(rows, err)
}
