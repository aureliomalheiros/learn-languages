package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

type Product struct {
	ID			int 	`gorm:"primarykey"`
	Name		string
	Price		float64
	CategoryID	int
	Category	Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID 		int `gorm:"primaryKey"`
	Name 	string
	Products []Product
}

type SerialNumber struct {
	ID int `gorm:"primarykey"`
	Number string
	ProductID int
}
func main(){
	dsm := "root:root@tcp(localhost:3306)/helloworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Eletronics"}
	db.Create(&category)

	db.Create(&Product{
		Name: "Mouse",
		Price: 40.00,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number: "12345456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}