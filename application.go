package main

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Number struct {
	Result int
}

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("ADDRESS"))
	if err != nil {
		panic("failed to connect database, error: " + err.Error())
	}
	err = db.CreateTable(&Number{}).Error
	if err != nil {
		panic("can't create a table numbers, error: " + err.Error())
	}
	return db
}

func WriteResultToTable(db *gorm.DB, result int) error {
	return db.Create(&Number{
		Result: result,
	}).Error
}

func GetResultFromTable(db *gorm.DB) (int, error) {
	var n []Number
	err := db.Find(&n).Error
	if err != nil {
		return 0, err
	}
	if len(n) == 0 {
		return 0, errors.New("can't find any number in database")
	}
	return n[0].Result, nil
}

func Plus(a int, b int) (result int) {
	result = a + b
	return
}
