package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_"fmt"
)

type Product struct {
	ID			int 	`gorm:"primarykey"`
	Name		string
	Price		float64
	CategoryID	int
	Category	Category
	gorm.Model
}

type Category struct {
	ID 		int `gorm:"primaryKey"`
	Name 	string
}
func main(){
	dsm := "root:root@tcp(localhost:3306)/helloworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Eletronics"}
	db.Create(&category)

	db.Create(&Product{
		Name: "TV",
		Price: 5000.00,
		CategoryID: category.ID,
	})

}