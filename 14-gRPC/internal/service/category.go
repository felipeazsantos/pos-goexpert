package service

import (
	"context"
	"io"

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

func (s *CategoryService) ListCategories(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.CategoryDB.GetAll()
	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoryList}, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.GetByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		category, err := s.CategoryDB.Create(req.GetName(), req.GetDescription())
		if err != nil {
			return err
		}

		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categories.Categories = append(categories.Categories, categoryResponse)
	}
}

func (s *CategoryService) CreateCategoryStreamBiDirectional(stream pb.CategoryService_CreateCategoryStreamBiDirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		category, err := s.CategoryDB.Create(req.GetName(), req.GetDescription())
		if err != nil {
			return err
		}

		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		if err := stream.Send(categoryResponse); err != nil {
			return err
		}
	}
}
