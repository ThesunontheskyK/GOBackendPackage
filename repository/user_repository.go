package repository

import (
	"encoding/json"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/thesunonthesky/GoBackendPackage/model"
)

type UserRepository interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
}

type UserRepositoryDB struct {
	db       *sqlx.DB
	mockUser []model.User
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	var users []model.User
	data, err := os.ReadFile("users.json")
	if err == nil {
		json.Unmarshal(data, &users)
	} else {
		return nil
	}
	return &UserRepositoryDB{
		db:       db,
		mockUser: users,
	}
}

func (r *UserRepositoryDB) Create(user model.User) error {
	query := `Insert Into users(name,email) values(?,?)`
	_, error := r.db.Exec(query, user.Name, user.Email)

	return error
}

func (r *UserRepositoryDB) GetAll() ([]model.User, error) {
	return r.mockUser, nil
}

func (r *UserRepositoryDB) GetByID(id int) (model.User, error) {
	for _, user := range r.mockUser {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, nil
}
