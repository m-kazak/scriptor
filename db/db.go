package db

import (
	"database/sql"
	"fmt"
	"errors"

	"scriptor/config"
	
	_"github.com/lib/pq"
)

type DataBase struct {
	// Database connector
	db *sql.DB

	// Database transaction
	tx *sql.Tx
}

var dataBase DataBase

// OpenTx is opening new transaction. If transaction wasn't closed throwing error
func (d *DataBase) OpenTx() (error) {
	transaction, err := d.db.Begin()
	if err != nil {
		return errors.New(err.Error())
	}
	d.tx = transaction
	return nil
}

// CommitTx is commiting transaction. If transaction cannot be commited - throwing error
func (d *DataBase) CommitTx() (error) {
	err := d.tx.Commit()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// RollbackTx is rollback transaction. If transaction cannot be rollbacked - throwing error
func (d *DataBase) RollbackTx() (error) {
	err := d.tx.Rollback()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

//ConnectToDB is used to connect to database and return database object
func ConnectDB() DataBase {
	if dataBase.db != nil {
		return dataBase
	}

	host := config.Config.Database.Host
	port := config.Config.Database.Port
	dbname := config.Config.Database.Dbname
	user := config.Config.Database.User
	password := config.Config.Database.Password

	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", params)

	if err != nil {
		config.Loggy.Error("Can't connect to DataBase")
	}
	dataBase.db = db

	return dataBase
}

//DisconnectDB is used to disconnect from database
func DisconnectDB() error {
	return dataBase.db.Close()
}
