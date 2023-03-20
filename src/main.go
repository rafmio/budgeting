package main

import (
	"fmt"
	"os"
)

func main() {
	menu()
}

func menu() {
	fmt.Println("Menu:")
	fmt.Printf("1. Show entries\n2. Add entry\n")

	var action int
	fmt.Printf("\nChoose action: ")
	fmt.Scanf("%d", &action)

	switch action {
	case 1:
		fmt.Println("Listing entries")
		entries, err := GetEntriesRows()
		if err != nil {
			fmt.Println("getting entries: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("in main.go file, case 1: ", entries)

	case 2:
		var entry Entry = Entry{
			Account:       "Alfa",
			Trdate:        "2023.02.24",
			Trtype:        "Purchases",
			Docdate:       "2023.02.25",
			Docnumb:       "43",
			Counterparty:  "Global Trucs",
			Cntp_tax_id:   "1660432383",
			Cntp_contract: "00942",
			Purpose:       "Truck repair, invoce #1343",
			Direction:     "outflow",
			Amount:        128_325.60,
			Item:          "Truck repair",
		}

		InsertEntry(entry)

	}
}
