package main

import "gorm.io/gorm"

func TransactionPost(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		post := Post{
			Title: "Hello World",
		}
		if err := tx.Create(&post).Error; err != nil {
			return err
		}
		comment := Comment{
			Content: "Hello World",
			PostID:  post.ID,
		}
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}
		return nil
	})
}

func TransactionPostWithManually(db *gorm.DB) error {
	tx := db.Begin() // 开启事物，并返回 tx

	post := Post{
		Title: "Hello World Manually",
	}
	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		return err
	}
	comment := Comment{
		Content: "Hello World Manually",
		PostID:  post.ID,
	}
	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
