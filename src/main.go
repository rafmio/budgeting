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

		fmt.Println(len(entries))
		for _, val := range entries {
			fmt.Println(val)
		}

	case 2:
		var entry Entry = Entry{
			Account:       "Alfa",
			Trdate:        "2023.02.26",
			Trtype:        "Purchases",
			Docdate:       "2023.02.26",
			Docnumb:       "543",
			Counterparty:  "Scania Rus",
			Cntp_tax_id:   "7854433454",
			Cntp_contract: "78434",
			Purpose:       "Truck repair, invoce #1232-32",
			Direction:     "outflow",
			Amount:        17_323.00,
			Item:          "Truck repair",
		}

		InsertEntry(entry)

	}
}
