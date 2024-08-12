package dtos

type CourseDTO struct {
	CourseId uint   `json:"course_id"`
	Title    string `json:"title"`
	Img      string `json:"img"`
}
