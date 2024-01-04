package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

type Product struct {
	ID		int 	`gorm:"primarykey"`
	Name	string
	Price	float64
}

func main(){
	dsm := "root:root@tcp(localhost:3306)/helloworld"
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }
	// db.Create(&products)

	//select one
	//var product Product
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	//select all
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}