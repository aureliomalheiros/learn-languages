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

	var p Product

	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	db.Delete(&p2)
}