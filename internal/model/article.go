package model

import "time"

type Article struct {
	ArticleId int `gorm:"primaryKey" json:"article_id"`
	SenderId int `gorm:"foreignKey" json:"sender_id"`
	Title string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
