package dto

type CreateCourseDTO struct {
	Title string `json:"title" binding:"required"`
	Img   string `json:"img" binding:"required"`
}
