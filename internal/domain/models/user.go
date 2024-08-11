package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:user_id"`
	CourseId      *uint     `gorm:"column:course_id"`
	Name          string    `gorm:"column:name"`
	ImgURL        string    `gorm:"column:img_url"`
	Username      string    `gorm:"unique;column:username"`
	Password      string    `gorm:"column:password"`
	Gender        string    `gorm:"column:gender"`
	Heart         int       `gorm:"column:heart;default:5"`
	LastHeartTime time.Time `gorm:"column:last_heart_time"`
	Points        int       `gorm:"column:points"`
	IsMembership  bool      `gorm:"column:is_membership;default:false"`
	IsAdmin       bool      `gorm:"column:is_admin;default:false"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	Course        Course    `gorm:"foreignKey:CourseId"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.UserId == "" {
		u.UserId = uuid.NewString()
	}
	return nil
}
