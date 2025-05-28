package model

type Key_word struct {
	KeyWordId int `gorm:"primaryKey" json:"key_word_id"`
	ArticleId int `gorm:"foreignKey" json:"article_id"`
	Content string `gorm:"not null" json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
