package controller

import "database/sql"

func Open() {
	db, err = sql.Open("mysql", "root:@/go_login_test")
	if err != nil {
		panic(err.Error())
	}

}
