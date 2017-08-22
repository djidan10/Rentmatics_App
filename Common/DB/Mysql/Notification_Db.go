package Mysql

import (
	Model "Rentmatics_App/Model"
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insertnotify_DB(Notify Model.Notification) {
	rows, err := OpenConnection["Rentmatics"].Exec("insert into notification (notifyname,notifyemail,notifyphone,notifytenant,notifyavail,notifycity) values (?,?,?,?,?,?)", Notify.Notifyname, Notify.NotifyEmail, Notify.Notifyphone, Notify.Notifytenenattype, Notify.Notifyavailibity, Notify.NotifyCity)
	fmt.Println("successfully inserted", rows, err)

}
