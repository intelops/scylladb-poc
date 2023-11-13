package services

import (
	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/daos"
	"github.com/MrAzharuddin/scylladb-gin/pkg/src/rest/models"
)

type UserService struct {
	userDao *daos.UserDao
}

func NewUserService() (*UserService, error) {
	userDao, err := daos.NewUserDao()
	if err != nil {
		return nil, err
	}
	return &UserService{userDao: userDao}, nil
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	return u.userDao.CreateUser(user)
}

func (u *UserService) GetUser(id string) (*models.User, error) {
	return u.userDao.GetUser(id)
}

func (u *UserService) GetUsers() ([]*models.User, error) {
	return u.userDao.GetUsers()
}

func (u *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return u.userDao.UpdateUser(user)
}

func (u *UserService) DeleteUser(id string) error {
	return u.userDao.DeleteUser(id)
}