package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ozonebg/gofluence/interfaces"
	"github.com/ozonebg/gofluence/models"
	"github.com/ozonebg/gofluence/repository"
	"github.com/ozonebg/gofluence/utils"
	log "github.com/sirupsen/logrus"
)

var usersLogger = log.WithField("component", "users controller")

type userController struct {
	usersRepository interfaces.UsersRepository
}

// NewUsersController returns a new controller for the articles.
func NewUsersController(usersRepo interfaces.UsersRepository) interfaces.UsersController {
	return &userController{
		usersRepository: usersRepo,
	}
}

func (uc *userController) AllUsers(w http.ResponseWriter, r *http.Request) {
	usersLogger.Info("endpoint hit: all users")

	users, err := uc.usersRepository.All()

	if err != nil {
		if err.Error() == repository.NotFoundArticlesError {
			usersLogger.Info("no users found")
			json.NewEncoder(w).Encode(users)
			return
		}
		usersLogger.WithError(err).Info("failed to get artiuserscles")
	}

	json.NewEncoder(w).Encode(users)
}

func (uc *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	usersLogger.Info("endpoint hit: get user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		usersLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	user, err := uc.usersRepository.GetUser(id)
	if err != nil {
		usersLogger.WithError(err).WithField("user_id", id).Error("failed to get user")
		fmt.Fprint(w, "no user found")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	usersLogger.Info("endpoint hit: create user")

	body, err := utils.ReadRequestBody(r)
	if err != nil {
		usersLogger.WithError(err).Info("failed to read body contents")
	}

	var user models.User
	json.Unmarshal(body, &user)

	err = uc.usersRepository.CreateUser(&user)
	if err != nil {
		usersLogger.WithError(err).Info("failed to create user")
	}
}

func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	usersLogger.Info("endpoint hit: update user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		usersLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	var updatedUser models.User
	body, err := utils.ReadRequestBody(r)
	if err != nil {
		usersLogger.WithError(err).Info("failed to read body contents")
	}

	json.Unmarshal(body, &updatedUser)

	err = uc.usersRepository.UpdateUser(id, &updatedUser)
	if err != nil {
		usersLogger.WithError(err).Info("failed to update user")
	}
}

func (uc *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	usersLogger.Info("endpoint hit: delete user")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		usersLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	err = uc.usersRepository.DeleteUser(id)
	if err != nil {
		usersLogger.WithError(err).Info("failed to delete user")
	}
}
