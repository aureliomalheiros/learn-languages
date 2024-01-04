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

	var products []Product
	db.Where("name Like ?", "%book%").Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}
}