package dtos

import "time"

type RegisterDTO struct {
	Name     string `json:"name"     binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordDTO struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type UpdateDTO struct {
	Name   string `json:"name"`
	Heart  int    `json:"heart"`
	Point  int    `json:"point"`
	Gender string `json:"gender"`
}

type UserDTO struct {
	UserId        string    `json:"user_id"`
	Name          string    `json:"name"`
	ImgURL        string    `json:"img_url"`
	Username      string    `json:"username"`
	Heart         int       `json:"heart"`
	LastHeartTime time.Time `json:"last_heart_time"`
	Points        int       `json:"point"`
	Gender        string    `json:"gender"`
	IsMembership  bool      `json:"is_membership"`
	IsAdmin       bool      `json:"is_admin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UploadDTO struct {
	Username string `json:"username"`
	ImgURL   string `json:"img_url"`
}

type LeaderboardDTO struct {
	Username string `json:"username"`
	ImgURL   string `json:"img_url"`
	Points   int    `json:"point"`
}
