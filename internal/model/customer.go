package model

import "time"

type Customer struct {
	CustomerID   int   `gorm:"primaryKey" json:"customer_id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Role      string    `gorm:"size:255;not null" json:"role"`
	StudyArea string    `gorm:"not null" json:"study_area"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
