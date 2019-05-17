package api

import (
	"context"

	"github.com/dgrijalva/jwt-go"

	"github.com/ozonebg/gofluence/internal/config"
	"github.com/ozonebg/gofluence/internal/models"

	"github.com/ozonebg/gofluence/internal/interfaces"
	"github.com/ozonebg/gofluence/pkg/api"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var usersLogger = log.WithField("component", "users service")

type userServiceServer struct {
	ur interfaces.UsersRepository
}

// NewUserServiceServer returns a newly created user service server.
func NewUserServiceServer(usersRepository interfaces.UsersRepository) api.UserServiceServer {
	return &userServiceServer{ur: usersRepository}
}

// Create new user.
func (s *userServiceServer) Create(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	usersLogger.Info("creating new user")

	usr := &api.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	}

	err := s.ur.CreateUser(usr)

	if err != nil {
		return nil, err
	}

	resp := &api.CreateUserResponse{
		Id:      usr.Id,
		Success: true,
	}

	return resp, nil
}

// AllUsers lists all users.
func (s *userServiceServer) All(ctx context.Context, req *api.AllUsersRequest) (*api.AllUsersResponse, error) {
	usersLogger.Info("getting all users")

	users, err := s.ur.All()
	if err != nil {
		return nil, err
	}

	resp := &api.AllUsersResponse{
		Users: users,
	}

	return resp, nil
}

// GetUser returns the asked user.
func (s *userServiceServer) Get(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	usersLogger.WithField("user_id", req.GetId()).Info("getting user")

	user, err := s.ur.GetUser(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	resp := &api.GetUserResponse{
		User: user,
	}

	return resp, nil
}

// Authenticate the user.
func (s *userServiceServer) Authenticate(ctx context.Context, req *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	usersLogger.WithField("username", req.GetUsername()).Info("trying to authenticate user")

	user, err := s.ur.GetUserByUsername(req.GetUsername())
	if err != nil {
		usersLogger.WithError(err).Info("failed to authenticate user")
	}

	if user == nil {
		resp := &api.AuthenticateResponse{
			Success: false,
			Error:   "no user information found",
		}
		return resp, nil
	}

	//Found user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		resp := &api.AuthenticateResponse{
			Success: false,
			Error:   "invalid login credentials. please try again",
		}
		return resp, nil
	}

	//Create jwt token
	tk := &models.Token{UserID: int(user.Id)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.GetTokenPassword()))

	resp := &api.AuthenticateResponse{
		Success: true,
		Token:   tokenString,
	}

	return resp, nil
}

// Update the user.
func (s *userServiceServer) Update(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	usersLogger.WithField("user_id", req.GetId()).Info("updating user")

	usr := &api.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
	}

	err := s.ur.UpdateUser(int(req.GetId()), usr)

	if err != nil {
		return nil, err
	}

	resp := &api.UpdateUserResponse{
		Id:      req.GetId(),
		Success: true,
	}

	return resp, nil
}

// Delete the user.
func (s *userServiceServer) Delete(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	usersLogger.WithField("user_id", req.GetId()).Info("deleting user")

	err := s.ur.DeleteUser(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	resp := &api.DeleteUserResponse{
		Success: true,
	}

	return resp, nil
}
