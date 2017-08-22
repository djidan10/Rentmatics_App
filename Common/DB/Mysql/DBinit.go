package Mysql

import (
	InitDb "Rentmatics_App/Common/DB/Mysql/InitializeDB"
	"database/sql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	OpenConnection = InitDb.Ret()
}
