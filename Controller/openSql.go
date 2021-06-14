package controller

import "database/sql"

func Open() {
	db, err = sql.Open("mysql", "root:@/gm_tool_test")
	if err != nil {
		panic(err.Error())
	}

}
