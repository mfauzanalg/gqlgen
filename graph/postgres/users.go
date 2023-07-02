package postgres

import (
	"golang/gqlgen/graph/models"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
