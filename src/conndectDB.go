package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type ConnectionConfig struct {
	DBMS        string // Database management system
	User        string // User of DBMS
	Password    string // Password to access
	NetLocation string // Net location of database: localhost ro IP
	Port        string // :port
	DBName      string // DB name
	Param       string // ?param1=value1&..
}

func (cs *ConnectionConfig) GetConnString() string {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s",
		cs.DBMS,
		cs.User,
		cs.Password,
		cs.NetLocation,
		cs.Port,
		cs.DBName,
		cs.Param,
	)

	return connString
}

// Hardcoded connection instance
var ConnInstance ConnectionConfig = ConnectionConfig{
	DBMS:        "postgresql",
	User:        "postgres",
	Password:    "qwq121",
	NetLocation: "194.58.102.129",
	Port:        "8432",
	DBName:      "budgeting",
	Param:       "sslmode=disable",
}

func EstablishConnectionDB() (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(ConnInstance.GetConnString())
	if err != nil {
		fmt.Println("parsing config: ", err.Error())
	} else {
		fmt.Println("parsing config: Successs")
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("connection to database: ", err.Error())
	} else {
		fmt.Println("connection to database: Success")
	}

	return conn, err
}
