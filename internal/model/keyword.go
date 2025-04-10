package model

import "time"

type Key_word struct {
	KeyWordId int `gorm:"primaryKey" json:"key_word_id"`
	ArticleId int `gorm:"foreignKey" json:"article_id"`
	Content string `gorm:"not null" json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
