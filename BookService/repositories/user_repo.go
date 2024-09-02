package repo

import (
	"book_service/configs"
	"book_service/models"
	pb "book_service/proto/user"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
)

type UserRepo struct {
	Config configs.ConfigurationsInterface
	Ctx    context.Context
}

func (r *UserRepo) Get(userId int) (models.User, error) {
	var user models.User

	// Create a gRPC connection
	address := r.Config.GRPCConfiguration().UserServiceGRPCAddress
	if address == "" {
		log.Printf("gRPC address is not configured properly")
		return user, errors.New("service address not configured")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to gRPC service at %s: %v", address, err)
		return user, errors.New("service unavailable")
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewUserServiceClient(conn) // Corrected client creation

	// Call ValidateToken gRPC method
	res, err2 := client.GetUserById(r.Ctx, &pb.GetUserByIdRequest{UserId: uint64(userId)})
	if err2 != nil {
		log.Printf("Failed : %v", err2)
		return user, err2
	}
	if res.Status == 0 {
		return user, errors.New(res.ErrorMessage)
	}

	// Convert res.User to models.User
	user = models.User{
		ID:        uint(res.User.Id),
		Username:  res.User.Username,
		Email:     res.User.Email,
		FirstName: res.User.FirstName,
		LastName:  res.User.LastName,
		Role:      res.User.Role,
		CreatedAt: res.User.CreatedAt,
		UpdatedAt: res.User.UpdatedAt,
		DeletedAt: res.User.DeletedAt,
	}

	return user, nil
}
