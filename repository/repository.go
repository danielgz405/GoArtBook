package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

type Repository interface {
	//Users
	InsertUser(ctx context.Context, user *models.InsertUser) (*models.Profile, error)
	GetUserById(ctx context.Context, id string) (*models.Profile, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, data models.UpdateUser) (*models.Profile, error)
	DeleteUser(ctx context.Context, id string) error

	//Close the connection
	Close() error
}

var implementation Repository

// Repo
func SetRepository(repository Repository) {
	implementation = repository
}

// Close the connection
func Close() error {
	return implementation.Close()
}
