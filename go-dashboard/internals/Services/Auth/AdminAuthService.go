package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	config "go-dashboard/internals/Config"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AdminAuthService struct {
	repo interfaces.RepositoryInterface[models.Admin]
}

func NewAdminAuthService(repo interfaces.RepositoryInterface[models.Admin]) *AdminAuthService {
	return &AdminAuthService{repo: repo}
}

func (s *AdminAuthService) Login(email string, password string) (string, error) {
	admins, err := s.repo.FindByWhere(map[string]interface{}{"email": email})
	if err != nil || len(admins) == 0 {
		return "", errors.New("invalid credentials")
	}

	admin := admins[0]

	hash := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	if admin.Password != hashedPassword {
		return "", errors.New("invalid credentials")
	}

	return s.generateToken(&admin)
}

func (s *AdminAuthService) generateToken(admin *models.Admin) (string, error) {
	expHours, _ := strconv.Atoi(config.AppConfig.JWTExpirationHours)
	if expHours == 0 {
		expHours = 24
	}

	claims := jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(expHours)).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func (s *AdminAuthService) ValidateToken(tokenString string) (*models.Admin, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid token sub claim")
	}

	admin, err := s.repo.FindById(uint(sub))
	if err != nil || admin == nil {
		return nil, errors.New("admin not found")
	}

	return admin, nil
}

func (s *AdminAuthService) RefreshToken(tokenString string) (string, error) {
	admin, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	return s.generateToken(admin)
}
