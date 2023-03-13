package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Entry struct {
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

func InsertEntry(conn *pgx.Conn, entry Entry) {
	query := fmt.Sprintf("INSERT INTO transactions (account, trdate, trtype, docdate, docnumb, counterpary, cntp_tax_id, cntp_contract, purpose, comnt, direction, amount, item) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%.2f', '%s')",
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

	_, err := conn.Exec(context.Background(), query)
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
