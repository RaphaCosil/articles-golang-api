package model

type User struct {
	userId int `gorm:"primaryKey" json:"user_id"`
	name String `gorm:"size:255;not null"`
	email String `gorm:"size:255;not null"`
	password String `gorm:"size:255;not null"`
	role String `gorm:"size:255;not null"`
	studyArea String `gorm:"not null" json:"study_area"`
	createdAt time.Time `gorm:"not null" json:"created_at"`
	updatedAt time.Time `gorm:"not null" json:"updated_at"`
}