package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

// GetProductByID retrieves a product by its ID using the repository.
// This Product was not supposed to be returned directly, but rather through a DTO object
func (uc *ProductUseCase) GetProductByID(id int64) (*Product, error) {
	return uc.repository.GetProductByID(id)
}