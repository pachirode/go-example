package main

import "gorm.io/gorm"

func SearchPost(db *gorm.DB) error {
	// 使用 Post 主键
	var post Post
	var comments []*Comment
	post.ID = 1

	// SELECT * FROM `comment` WHERE `comment`.`post_id` = 1 AND `comment`.`deleted_at` IS NULL
	err := db.Model(&post).Association("Comments").Find(&comments)
	if err != nil {
		return err
	}

	// 预加载
	post2 := Post{}
	err = db.Preload("Comments").Preload("Tags").First(post2).Error
	if err != nil {
		return err
	}

	return nil
}
