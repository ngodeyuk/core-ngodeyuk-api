package services

import (
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
)

type CourseService interface {
	Create(dto *dtos.CourseDTO) error
}

type courseService struct {
	repository repositories.CourseRepository
}

func NewCourseService(repository repositories.CourseRepository) CourseService {
	return &courseService{repository}
}

func (service *courseService) Create(dto *dtos.CourseDTO) error {
	course := &models.Course{
		Title: dto.Title,
		Img:   dto.Img,
	}

	return service.repository.Create(course)
}
