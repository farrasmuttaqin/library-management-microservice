package usecase

import (
	"context"
	"time"
	"user_service/proto/user"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
	UserUseCase *UserUseCase
}

// ValidateToken handles the token validation gRPC request
func (s *UserServiceServer) ValidateToken(ctx context.Context, req *user.ValidateTokenRequest) (*user.ValidateTokenResponse, error) {
	// Validate the token
	claims, err := s.UserUseCase.ValidateToken(req.GetToken())
	if err != nil {
		return &user.ValidateTokenResponse{
			Status:       0, // Failure
			ErrorMessage: err.Error(),
		}, nil
	}

	// Return a successful response with claims
	return &user.ValidateTokenResponse{
		Status: 1, // Success
		Claims: &user.Claims{
			UserId:   uint64(claims.UserID),
			Username: claims.Username,
			Role:     claims.Role,
			Exp:      claims.ExpiresAt.Unix(),
			Iat:      claims.IssuedAt.Unix(),
		},
	}, nil
}
func (s *UserServiceServer) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	// Retrieve user details by ID using the UserUseCase
	userDetails, err := s.UserUseCase.GetByUserId(uint(req.GetUserId()))
	if err != nil {
		return &user.GetUserByIdResponse{
			Status:       0, // Failure
			ErrorMessage: err.Error(),
		}, nil
	}

	// Return a successful response with user details
	return &user.GetUserByIdResponse{
		Status: 1, // Success
		User: &user.User{
			Id:        uint64(userDetails.ID),
			Username:  userDetails.Username,
			Email:     userDetails.Email,
			FirstName: userDetails.FirstName,
			LastName:  userDetails.LastName,
			Role:      userDetails.Role,
			CreatedAt: userDetails.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: userDetails.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt: formatNullableTime(&userDetails.DeletedAt.Time),
		},
	}, nil
}

// Helper function to format nullable time
func formatNullableTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
