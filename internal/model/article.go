package model

// import "time"

type Article struct {
	ArticleId int `gorm:"primaryKey" json:"article_id"`
	SenderId int `gorm:"foreignKey" json:"sender_id"`
	Title string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
