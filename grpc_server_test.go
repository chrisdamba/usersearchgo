package main

import (
	"context"
	"github.com/chrisdamba/usersearchgo/services"
	"testing"

	pb "github.com/chrisdamba/usersearchgo/proto"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	var server = &UserServiceServer{
		service: services.NewUserService(),
	}
	req := &pb.UserIdRequest{Id: 1}
	resp, err := server.GetUserById(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.User.Id)
}

func TestGetUsersByIds(t *testing.T) {
	server := &UserServiceServer{
		service: services.NewUserService(),
	}

	req := &pb.UserListRequest{Ids: []int32{1, 2}}
	resp, err := server.GetUsersByIds(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Users, 2)
}

func TestSearchUsers(t *testing.T) {
	server := &UserServiceServer{
		service: services.NewUserService(),
	}

	req := &pb.UserSearchRequest{Fname: "John"}
	resp, err := server.SearchUsers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	for _, user := range resp.Users {
		assert.Equal(t, req.Fname, user.Fname)
	}
}

func TestSearchUsersByMultipleCriteria(t *testing.T) {
	server := &UserServiceServer{
		service: services.NewUserService(),
	}

	req := &pb.UserSearchRequest{City: "New York", Married: true}
	resp, err := server.SearchUsers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	for _, user := range resp.Users {
		assert.Equal(t, req.City, user.City)
		assert.Equal(t, req.Married, user.Married)
	}
}
