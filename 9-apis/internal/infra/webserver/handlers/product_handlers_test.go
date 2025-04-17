package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductDB struct {
	mock.Mock
}

func (m *MockProductDB) Create(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	args := m.Called(page, limit, sort)
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *MockProductDB) FindByID(id string) (*entity.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *MockProductDB) Update(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductDB) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("should create product successfully", func(t *testing.T) {
		product := entity.Product{
			Name:        "Test Product",
			Price:       10.0,
			Description: "Test Description",
		}
		mockDB.On("Create", mock.AnythingOfType("*entity.Product")).Return(nil)

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
		w := httptest.NewRecorder()

		handler.CreateProduct(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response entity.Product
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.ID)
		mockDB.AssertExpectations(t)
	})

	t.Run("should return error with invalid json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		handler.CreateProduct(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestListProducts(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("should list products successfully", func(t *testing.T) {
		products := []entity.Product{
			{
				ID:          uuid.New(),
				Name:        "Product 1",
				Price:       10.0,
				Description: "Description 1",
			},
			{
				ID:          uuid.New(),
				Name:        "Product 2",
				Price:       20.0,
				Description: "Description 2",
			},
		}
		mockDB.On("FindAll", 1, 10, "asc").Return(products, nil)

		req := httptest.NewRequest(http.MethodGet, "/products?page=1&limit=10&sort=asc", nil)
		w := httptest.NewRecorder()

		handler.ListProducts(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response []entity.Product
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		mockDB.AssertExpectations(t)
	})
}

func TestGetProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("should get product successfully", func(t *testing.T) {
		id := uuid.New()
		product := &entity.Product{
			ID:          id,
			Name:        "Test Product",
			Price:       10.0,
			Description: "Test Description",
		}
		mockDB.On("FindByID", id.String()).Return(product, nil)

		req := httptest.NewRequest(http.MethodGet, "/products/"+id.String(), nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id.String())
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.GetProduct(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response entity.Product
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, id, response.ID)
		mockDB.AssertExpectations(t)
	})

	t.Run("should return error when id is empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/", nil)
		rctx := chi.NewRouteContext()
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.GetProduct(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("should update product successfully", func(t *testing.T) {
		id := uuid.New()
		product := entity.Product{
			ID:          id,
			Name:        "Updated Product",
			Price:       20.0,
			Description: "Updated Description",
		}
		mockDB.On("Update", &product).Return(nil)

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPut, "/products/"+id.String(), bytes.NewReader(body))
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id.String())
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.UpdateProduct(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response entity.Product
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, id, response.ID)
		mockDB.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("should delete product successfully", func(t *testing.T) {
		id := uuid.New()
		mockDB.On("Delete", id.String()).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/products/"+id.String(), nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id.String())
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.DeleteProduct(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockDB.AssertExpectations(t)
	})

	t.Run("should return error when id is empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/", nil)
		rctx := chi.NewRouteContext()
		req = req.WithContext(context.WithValue(context.Background(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.DeleteProduct(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
