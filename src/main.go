package main

import (
	"fmt"
	"os"
)

func main() {
	var entry Entry = Entry{
		Account:       "Sber",
		Trdate:        "2023.02.02",
		Trtype:        "Loans",
		Docdate:       "2023.02.02",
		Docnumb:       "509/D23",
		Counterparty:  "Sber-Bank",
		Cntp_tax_id:   "7759435433",
		Cntp_contract: "LS0232",
		Purpose:       "Payment #11 for contract LS0232",
		Direction:     "outflow",
		Amount:        130430.58,
		Item:          "Credits and loans",
	}

	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB: ", err.Error())
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	InsertEntry(conn, entry)
}
