package services

import (
	"ngodeyuk-core/internal/domain/dtos"
	"ngodeyuk-core/internal/domain/models"
	"ngodeyuk-core/internal/domain/repositories"
)

type CourseService interface {
	Create(dto *dtos.CourseDTO) error
	Update(courseId uint, dto *dtos.CourseDTO) error
	GetAll() ([]models.Course, error)
	GetByID(courseId uint) (*models.Course, error)
	DeleteByID(courseId uint) error
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

func (service *courseService) Update(courseId uint, dto *dtos.CourseDTO) error {
	course, err := service.repository.FindByID(courseId)
	if err != nil {
		return err
	}
	if dto.Title != "" {
		course.Title = dto.Title
	}
	if dto.Img != "" {
		course.Img = dto.Img
	}
	if err := service.repository.Update(course); err != nil {
		return err
	}
	return nil
}

func (service *courseService) GetAll() ([]models.Course, error) {
	return service.repository.FindAll()
}

func (service *courseService) GetByID(courseId uint) (*models.Course, error) {
	return service.repository.FindByID(courseId)
}

func (service *courseService) DeleteByID(courseId uint) error {
	course, err := service.repository.FindByID(courseId)
	if err != nil {
		return err
	}
	return service.repository.Delete(course)
}
