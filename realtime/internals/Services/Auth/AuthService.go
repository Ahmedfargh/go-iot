package auth

type AuthService struct {
}

func (ath *AuthService) ValidateToken(token string) (bool, error) {

	return true, nil
}
