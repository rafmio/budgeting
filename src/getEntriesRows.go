package main

import (
	"context"
	"fmt"
)

func GetEntriesRows() ([]string, error) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB: ", err.Error())
	} else {
		fmt.Println("connection to DB: Success")
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from transactions;")
	if err != nil {
		fmt.Println("extracting entries: ", err.Error())
	} else {
		fmt.Println("extracting entries: Success")
	}
	defer rows.Close()

	var sqlResponse []string

	for rows.Next() {
		var entry string
		rows.Scan(&entry)
		fmt.Println("in getEntriesRows.go file: ", entry)
		sqlResponse = append(sqlResponse, entry)
	}
	fmt.Println("slice with entries is completed")

	return sqlResponse, err
}
