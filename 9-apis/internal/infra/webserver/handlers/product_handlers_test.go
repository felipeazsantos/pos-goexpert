package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"errors"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	entityPkg "github.com/felipeazsantos/pos-goexpert/apis/pkg/entity"
	"github.com/go-chi/chi/v5"
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

	t.Run("success", func(t *testing.T) {
		product := &entity.Product{
			Name:        "Test Product",
			Price:       100.0,
			Description: "Test Description",
		}
		mockDB.On("Create", mock.AnythingOfType("*entity.Product")).Return(nil)

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
		w := httptest.NewRecorder()

		handler.CreateProduct(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockDB.AssertExpectations(t)
	})

	t.Run("invalid json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		handler.CreateProduct(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestListProducts(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	products := []entity.Product{
		{Name: "Product 1", Price: 100.0, Description: "Description 1"},
		{Name: "Product 2", Price: 200.0, Description: "Description 2"},
	}

	mockDB.On("FindAll", 0, 10, "").Return(products, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	handler.ListProducts(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

	var response []entity.Product
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
}

func TestGetProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("success", func(t *testing.T) {
		product := &entity.Product{
			Name:        "Test Product",
			Price:       100.0,
			Description: "Test Description",
		}
		mockDB.On("FindByID", "1").Return(product, nil)

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.GetProduct(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockDB.AssertExpectations(t)

		var response entity.Product
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, product.Name, response.Name)
	})

	t.Run("not found", func(t *testing.T) {
		mockDB.On("FindByID", "999").Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "999")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.GetProduct(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockDB.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("success", func(t *testing.T) {
		id := entityPkg.NewID()
		product := &entity.Product{
			ID:          id,
			Name:        "Updated Product",
			Price:       150.0,
			Description: "Updated Description",
		}
		
		// Mock FindByID first
		mockDB.On("FindByID", id.String()).Return(product, nil)
		// Then mock Update
		mockDB.On("Update", mock.AnythingOfType("*entity.Product")).Return(nil)

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(body))
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id.String())
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.UpdateProduct(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockDB.AssertExpectations(t)
	})

	t.Run("invalid json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader([]byte("invalid json")))
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.UpdateProduct(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	mockDB := new(MockProductDB)
	handler := NewProductHandler(mockDB)

	t.Run("success", func(t *testing.T) {
		// Mock FindByID first
		mockDB.On("FindByID", "1").Return(&entity.Product{}, nil)
		// Then mock Delete
		mockDB.On("Delete", "1").Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.DeleteProduct(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockDB.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockDB.On("FindByID", "999").Return(nil, errors.New("not found"))

		req := httptest.NewRequest(http.MethodDelete, "/products/999", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "999")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		w := httptest.NewRecorder()

		handler.DeleteProduct(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockDB.AssertExpectations(t)
	})
}
