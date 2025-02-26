package handlers

import (
    "context"
    "user-service/models"
    "user-service/db"
    "user-service/auth"

    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
    pb "user-service/proto"
)

type UserServiceServer struct {
    pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := models.User{
        ID:       uuid.New().String(),
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    if result := db.DB.Create(&user); result.Error != nil {
        return nil, result.Error
    }

    return &pb.RegisterResponse{UserId: user.ID, Message: "User registered successfully"}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    var user models.User
    if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        return nil, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return nil, err
    }

    token, err := auth.GenerateToken(user.ID)
    if err != nil {
        return nil, err
    }

    return &pb.LoginResponse{Token: token}, nil
}

func (s *UserServiceServer) GetUserProfile(ctx context.Context, req *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {
    var user models.User
    if err := db.DB.First(&user, "id = ?", req.UserId).Error; err != nil {
        return nil, err
    }

    return &pb.UserProfileResponse{UserId: user.ID, Email: user.Email, Name: user.Name}, nil
}
