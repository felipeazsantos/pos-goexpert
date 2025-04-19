package dto


type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}

type CreateUserInput struct {
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Password string  `json:"password"`
}

