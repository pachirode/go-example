package main

import "gorm.io/gorm"

type Animal struct {
	Name string
}

func CreateAnimals(db *gorm.DB) error {
	tx := db.Begin()
	defer tx.Rollback()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
