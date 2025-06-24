package usecase

import (
	"context"

	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/entity"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/internal/repository"
	"github.com/felipeazsantos/pos-goexpert/18-UOW/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseInterfaceUow interface {
	Execute(ctx context.Context, input InputUseCase) error
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(u uow.UowInterface) AddCourseUseCaseInterfaceUow {
	return &AddCourseUseCaseUow{
		Uow: u,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow uow.UowInterface) error {
		categoryRepo := a.getCategoryRepository(ctx)
		courseRepo := a.getCourseRepository(ctx)

		category := entity.Category{
			Name: input.CategoryName,
		}
		
		if err := categoryRepo.Insert(ctx, category); err != nil {
			return err
		}
	
		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
	
		if err := courseRepo.Insert(ctx, course); err != nil {
			return err
		}

		return nil
	})	
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context,) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(*repository.CategoryRepository)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context,) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(*repository.CourseRepository)
}

