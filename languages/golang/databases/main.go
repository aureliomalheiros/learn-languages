package main

import (
	"github.com/google/uuid"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID	string
	Name	string
	Price 	float64
}

func NewProduct(name string, price float64) *Product{
	return  &Product{
		ID: uuid.New().String(), 
		Name: name,
		Price: price,
	}
}

func insertProdut(db *sql.DB, product Product) error{
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil

}

func main(){
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/helloworld")
	if err != nil{
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	err = insertProdut(db, *product)
	if err != nil {
		panic(err)
	}
}