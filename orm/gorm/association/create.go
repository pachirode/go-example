package main

import "gorm.io/gorm"

func CreatePost(db *gorm.DB) error {
	var post Post
	post = Post{
		Title:   "post",
		Content: "content",
		Comments: []*Comment{
			{Content: "comment", Post: &post},
		},
		Tags: []*Tag{
			{Name: "tag"},
		},
	}

	result := db.Create(&post)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
