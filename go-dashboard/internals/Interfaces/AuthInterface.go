package interfaces

import (
	models "go-dashboard/internals/Models"
)

type AuthInterface interface {
	Login(email string, password string) (string, error)
	ValidateToken(token string) (*models.Admin, error)
	RefreshToken(token string) (string, error)
}
