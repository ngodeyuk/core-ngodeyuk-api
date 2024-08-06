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
	Heart         int       `gorm:"column:heart;default:5"`
	LastHeartTime time.Time `gorm:"column:last_heart_time"`
	Points        int       `gorm:"column:points"`
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
