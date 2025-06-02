package service

import (
	"context"

	"github.com/felipeazsantos/pos-goexpert/14-gRPC/internal/database"
	"github.com/felipeazsantos/pos-goexpert/14-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.CategoryDB.Create(req.GetName(), req.GetDescription())
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
