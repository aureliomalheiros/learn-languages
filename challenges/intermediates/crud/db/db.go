package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver conection MySQL
)

//Connection function
func Connect() (*sql.DB, error){
	stringConection := "golang:golang@tcp(172.17.0.2:3306)/golang?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConection)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}