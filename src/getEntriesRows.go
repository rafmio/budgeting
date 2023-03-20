package main

import (
	"context"
	"fmt"
)

func GetEntriesRows() ([]Entry, error) {
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

	var sqlResponse []Entry

	for rows.Next() {
		values, err := rows.Values()
		fmt.Println("values: ", values[11])
		if err != nil {
			fmt.Println("values: ", values)
		}
	}
	fmt.Println("-------------------")

	// for rows.Next() {
	// 	var entry Entry
	// 	rows.Scan(&entry.ID,
	// 		&entry.Account,
	// 		&entry.Trdate,
	// 		&entry.Trtype,
	// 		&entry.Docdate,
	// 		&entry.Docnumb,
	// 		&entry.Counterparty,
	// 		&entry.Cntp_tax_id,
	// 		&entry.Cntp_contract,
	// 		&entry.Purpose,
	// 		&entry.Comment,
	// 		&entry.Direction,
	// 		&entry.Amount,
	// 		&entry.Item,
	// 	)
	// 	fmt.Println("in getEntriesRows.go file: ", entry)
	// 	sqlResponse = append(sqlResponse, entry)
	// }
	// fmt.Println("slice with entries is completed")

	return sqlResponse, err
}
