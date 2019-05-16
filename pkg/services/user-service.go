package api

import (
	"context"
	"errors"

	"github.com/ozonebg/gofluence/pkg/api"
	log "github.com/sirupsen/logrus"
)

var usersLogger = log.WithField("component", "users service")

type userServiceServer struct {
}

// NewUserServiceServer returns a newly created user service server.
func NewUserServiceServer() api.UserServiceServer {
	return &userServiceServer{}
}

// Create new user.
func (s *userServiceServer) Create(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.WithField("request", *req).Info("create endpoint hit")

	return nil, errors.New("not implemented yet")
}

// AllUsers lists all users.
func (s *userServiceServer) All(ctx context.Context, req *api.AllUsersRequest) (*api.AllUsersResponse, error) {
	log.WithField("request", *req).Info("all endpoint hit")

	return nil, errors.New("not implemented yet")
}

// GetUser returns the asked user.
func (s *userServiceServer) Get(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	log.WithField("request", *req).Info("get endpoint hit")

	return nil, errors.New("not implemented yet")
}

// Authenticate the user.
func (s *userServiceServer) Authenticate(ctx context.Context, req *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	log.WithField("request", *req).Info("authenticate endpoint hit")

	return nil, errors.New("not implemented yet")
}

// Update the user.
func (s *userServiceServer) Update(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	log.WithField("request", *req).Info("authenticate endpoint hit")

	return nil, errors.New("not implemented yet")
}

// Delete the user.
func (s *userServiceServer) Delete(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	log.WithField("request", *req).Info("authenticate endpoint hit")

	return nil, errors.New("not implemented yet")
}
