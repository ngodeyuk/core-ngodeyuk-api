package services

import (
	"ngodeyuk-core/pkg/dto"
	"ngodeyuk-core/pkg/models"

	"gorm.io/gorm"
)

type CourseService interface {
	CreateCourse(dto.CreateCourseDTO) (models.Course, error)
}

type courseService struct {
	db *gorm.DB
}

func NewCourseService(db *gorm.DB) CourseService {
	return &courseService{db}
}

func (s *courseService) CreateCourse(courseDTO dto.CreateCourseDTO) (models.Course, error) {
	course := models.Course{
		Title: courseDTO.Title,
		Img:   courseDTO.Img,
	}
	err := s.db.Create(&course).Error
	return course, err
}
