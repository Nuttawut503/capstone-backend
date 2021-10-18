package graph

import "github.com/Nuttawut503/capstone-backend/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *db.PrismaClient
}
