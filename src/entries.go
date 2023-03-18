package main

import (
	"context"
	"fmt"
	"os"
)

type Entry struct {
	ID            int32
	Account       string
	Trdate        string
	Trtype        string
	Docdate       string
	Docnumb       string
	Counterparty  string
	Cntp_tax_id   string
	Cntp_contract string
	Purpose       string
	Comment       string
	Direction     string
	Amount        float64
	Item          string
}

var DBColumns map[string]string = map[string]string{
	"ID":            "transaction_id",
	"Account":       "account",
	"Trdate":        "trdate",
	"Trtype":        "trtype",
	"Docdate":       "docdate",
	"Docnumb":       "docnumb",
	"Counterparty":  "counterpary",
	"Cntp_tax_id":   "cntp_tax_id",
	"Cntp_contract": "cntp_contract",
	"Purpose":       "purpose",
	"Comment":       "comnt",
	"Direction":     "direction",
	"Amount":        "amount",
	"Item":          "item",
}

func InsertEntry(entry Entry) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB: ", err.Error())
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	query := fmt.Sprintf("INSERT INTO transactions (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%.2f', '%s')",
		DBColumns["Account"],
		DBColumns["Trdate"],
		DBColumns["Trtype"],
		DBColumns["Docdate"],
		DBColumns["Docnumb"],
		DBColumns["Counterparty"],
		DBColumns["Cntp_tax_id"],
		DBColumns["Cntp_contract"],
		DBColumns["Purpose"],
		DBColumns["Comment"],
		DBColumns["Direction"],
		DBColumns["Amount"],
		DBColumns["Item"],
		entry.Account,
		entry.Trdate,
		entry.Trtype,
		entry.Docdate,
		entry.Docnumb,
		entry.Counterparty,
		entry.Cntp_tax_id,
		entry.Cntp_contract,
		entry.Purpose,
		entry.Comment,
		entry.Direction,
		entry.Amount,
		entry.Item,
	)

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("adding entry: ", err.Error())
	} else {
		fmt.Println("adding entry: Success")
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	} else {
		fmt.Println("closing connection: Success")
	}

	defer conn.Close(context.Background())
}

func UpdateEntry(entry Entry) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB: ", err.Error())
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	query := fmt.Sprintf("INSERT INTO transactions (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%.2f', '%s')",
		DBColumns["Account"],
		DBColumns["Trdate"],
		DBColumns["Trtype"],
		DBColumns["Docdate"],
		DBColumns["Docnumb"],
		DBColumns["Counterparty"],
		DBColumns["Cntp_tax_id"],
		DBColumns["Cntp_contract"],
		DBColumns["Purpose"],
		DBColumns["Comment"],
		DBColumns["Direction"],
		DBColumns["Amount"],
		DBColumns["Item"],
		entry.Account,
		entry.Trdate,
		entry.Trtype,
		entry.Docdate,
		entry.Docnumb,
		entry.Counterparty,
		entry.Cntp_tax_id,
		entry.Cntp_contract,
		entry.Purpose,
		entry.Comment,
		entry.Direction,
		entry.Amount,
		entry.Item,
	)

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("adding entry: ", err.Error())
	} else {
		fmt.Println("adding entry: Success")
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	} else {
		fmt.Println("closing connection: Success")
	}

	defer conn.Close(context.Background())
}
