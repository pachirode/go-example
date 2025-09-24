package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string     `gorm:"column:title"`
	Content  string     `gorm:"column:content"`
	Comments []*Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
	Tags     []*Tag     `gorm:"many2many:post_tags"`
}

func (p *Post) TableName() string {
	return "post"
}

type Comment struct {
	gorm.Model
	Content string `gorm:"column:content"`
	PostID  uint   `gorm:"column:post_id"`
	Post    *Post
}

func (c *Comment) TableName() string {
	return "comment"
}

type Tag struct {
	gorm.Model
	Name string  `gorm:"column:name"`
	Post []*Post `gorm:"many2many:post_tags"`
}

func (t *Tag) TableName() string {
	return "tag"
}
