package graph

import "github.com/felipeazsantos/pos-goexpert/13-graphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CategoryDB *database.Category
	CourseDB *database.Course
}
