package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

func main() {
	connectString := "sys/Oradoc_db1@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=localhost)(PORT=32769)))(CONNECT_DATA=(SERVICE_NAME=ORCLCDB.localdomain))) as sysdba"
	db, err := sql.Open("goracle", connectString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
