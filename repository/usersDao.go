package repository

import (
	"errors"

	"github.com/gocraft/dbr"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/ozonebg/gofluence/interfaces"
	"github.com/ozonebg/gofluence/models"
)

const (
	usersTableName = "users"

	// NotFoundUsersError is error if no User is found.
	NotFoundUsersError = "no users found"
)

var usersRepositoryLogger = logrus.WithField("component", "users dao")

type usersDao struct {
	s *dbr.Session
}

// NewUsersDao returns new UsersRepository
func NewUsersDao(session *dbr.Session) interfaces.UsersRepository {
	return &usersDao{
		s: session,
	}
}

func (a *usersDao) All() (models.Users, error) {
	tx, err := a.s.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	var users models.Users
	tx.Select("*").From(usersTableName).Load(&users)

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if users == nil {
		return models.Users{}, errors.New(NotFoundUsersError)
	}

	return users, nil
}

func (a *usersDao) GetUser(id int) (*models.User, error) {
	tx, err := a.s.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	var user models.User
	result, err := tx.Select("*").From(usersTableName).Where("id = ?", id).Load(&user)
	if err != nil {
		return nil, err
	}

	if result != 1 {
		return nil, errors.New(NotFoundUsersError)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *usersDao) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New(InvalidDataError)
	}
	tx, err := a.s.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPass)

	_, err = tx.InsertInto(usersTableName).Columns("username", "password", "email").Record(user).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *usersDao) DeleteUser(id int) error {
	tx, err := a.s.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.DeleteFrom(usersTableName).Where("id = ?", id).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *usersDao) UpdateUser(id int, updatedUser *models.User) error {
	tx, err := a.s.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	updateMap := getUpdateMapForUser(updatedUser)

	_, err = tx.Update(usersTableName).Where("id = ?", id).SetMap(updateMap).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func getUpdateMapForUser(user *models.User) map[string]interface{} {
	updateMap := make(map[string]interface{}, 3)
	updateMap["username"] = user.Username
	updateMap["password"] = user.Password
	updateMap["email"] = user.Email

	return updateMap
}

func (a *usersDao) GetUserByUsername(username string) (*models.User, error) {
	tx, err := a.s.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	var user models.User
	result, err := tx.Select("*").From(usersTableName).Where("username = ?", username).Load(&user)
	if err != nil {
		return nil, err
	}

	if result != 1 {
		return nil, errors.New(NotFoundUsersError)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
