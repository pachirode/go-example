package main

import "gorm.io/gorm"

func UpdatePost(db *gorm.DB) error {
	var post Post
	// SELECT * FROM `post` WHERE `post`.`deleted_at` IS NULL ORDER BY `post`.`id` LIMIT 1
	result := db.First(&post)
	if err := result.Error; err != nil {
		return err
	}

	// 替换关联
	comment := Comment{
		Content: "comment3",
	}
	err := db.Model(&post).Association("Comments").Replace([]*Comment{&comment})
	if err != nil {
		return err
	}

	return nil
}
