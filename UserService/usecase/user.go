package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user_service/configs"
	"user_service/helpers"
	"user_service/helpers/database"
	"user_service/models"
	repo "user_service/repositories"
)

type UserUseCase struct {
	Configs     configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	UserRepo    repo.UserRepository
}

func (u *UserUseCase) Login(loginData models.LoginForm) (map[string]interface{}, error) {
	getByUserName := map[string]interface{}{
		"username": loginData.UserName,
	}

	// Retrieve user by username
	user, err := u.UserRepo.Get(getByUserName)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to retrieve user",
		}, err
	}

	if user == nil {
		return map[string]interface{}{
			"status":  0,
			"message": "User not found",
		}, errors.New("user not found")
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Invalid password",
		}, errors.New("invalid password")
	}

	// Generate JWT token
	token, err2 := u.generateToken(user)
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to generate token",
		}, err2
	}

	return map[string]interface{}{
		"status":  1,
		"message": "success",
		"user":    user,
		"token":   token,
	}, nil
}

func (u *UserUseCase) Register(registerData models.RegisterForm) (map[string]interface{}, error) {
	// Validate input
	if err := u.validateRegisterForm(registerData); err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to hash password",
		}, err
	}

	// Create a new user
	user := &models.User{
		Username:  registerData.Username,
		Email:     registerData.Email,
		Password:  string(hashedPassword),
		FirstName: registerData.FirstName,
		LastName:  registerData.LastName,
		Role:      registerData.Role,
	}

	// Save user to the repository
	err = u.UserRepo.Create(user)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to create user",
		}, err
	}

	return map[string]interface{}{
		"status":  1,
		"message": "Registration successful",
		"user":    user,
	}, nil
}

func (u *UserUseCase) generateToken(user *models.User) (string, error) {
	claims := &models.JWTTokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(u.Configs.JWTConfiguration().ExpiryHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.Configs.JWTConfiguration().Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (u *UserUseCase) ValidateToken(jwtToken string) (*models.JWTTokenClaims, error) {
	// Retrieve the JWT secret from environment variables or configuration
	jwtSecret := u.Configs.JWTConfiguration().Secret
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(jwtToken, &models.JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's algorithm matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if claims, ok := token.Claims.(*models.JWTTokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetByUserId retrieves a user by their ID
func (u *UserUseCase) GetByUserId(userId uint) (*models.User, error) {
	// Retrieve the user by ID from the repository
	qUser := map[string]interface{}{
		"id": userId,
	}
	user, err := u.UserRepo.Get(qUser)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if helpers.IsEmptyStruct(user) {
		return nil, errors.New("user not found")
	}

	return user, nil
}
