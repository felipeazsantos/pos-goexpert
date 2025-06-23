package usecase

import (
	"context"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseInterface interface {
	Execute(ctx context.Context, input InputUseCase) error
}

type AddCourseUseCase struct {
	CategoryRepository repository.CategoryRepositoryInterface
	CourseRepository   repository.CourseRepositoryInterface
}

func NewAddCourseUseCase(categoryRepository repository.CategoryRepositoryInterface, courseRepository repository.CourseRepositoryInterface) AddCourseUseCaseInterface {
	return &AddCourseUseCase{
		CategoryRepository: categoryRepository,
		CourseRepository:   courseRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}
	
	if err := a.CategoryRepository.Insert(ctx, category); err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	if err := a.CourseRepository.Insert(ctx, course); err != nil {
		return err
	}

	return nil
}
