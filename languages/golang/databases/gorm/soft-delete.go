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
	gorm.Model
}

func main(){
	dsm := "root:root@tcp(localhost:3306)/helloworld?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name: "Notebook",
	// 	Price: 1000.00,
	// })
	
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	var p Product
	db.First(&p, 1)
	fmt.Println(p.Name)
	db.Delete(&p)
}