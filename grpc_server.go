package main

import (
	"context"
	pb "github.com/chrisdamba/usersearchgo/proto"
	"github.com/chrisdamba/usersearchgo/services"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	service *services.UserService
}

func (s *UserServiceServer) GetUserById(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	user := s.service.GetUserByID(int32(req.Id))
	if user == nil {
		return nil, nil
	}
	return &pb.UserResponse{
		User: &pb.User{
			Id:      int32(user.Id),
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		},
	}, nil
}

func (s *UserServiceServer) GetUsersByIds(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	users := s.service.GetUsersByIDs(req.Ids)
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      int32(user.Id),
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UserListResponse{Users: pbUsers}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.UserSearchRequest) (*pb.UserListResponse, error) {
	users := s.service.SearchUsers(req.Fname, req.City, req.Phone, req.Married)
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      int32(user.Id),
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UserListResponse{Users: pbUsers}, nil
}

func intSliceToInt32(slice []int32) []int {
	var result []int
	for _, v := range slice {
		result = append(result, int(v))
	}
	return result
}
