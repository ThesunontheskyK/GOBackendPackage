package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/thesunonthesky/GoBackendPackage/model"
	"github.com/thesunonthesky/GoBackendPackage/errs"
)

type UserRepository interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
	GetByID(id int) (model.User, error)
}

type UserRepositoryDB struct {
	db       *sqlx.DB
	mockUser []model.User
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	var users []model.User
	data, err := os.ReadFile("../../users.json")
	if err == nil {
		json.Unmarshal(data, &users)
	} else {
		fmt.Println("Warning: Could not read users.json:", err)
	}
	return &UserRepositoryDB{
		db:       db,
		mockUser: users,
	}
}

func (r *UserRepositoryDB) Create(user model.User) error {
		// เช็คอีเมลซ้ำ
		for _, u := range r.mockUser {
			if u.Email == user.Email {
				return errs.NewConflictError("Email already exists")
			}
		}

		maxID := 0
		for _,u:=range r.mockUser{
			if u.ID > maxID{
				maxID = u.ID
			}
		}
		
		user.ID = maxID + 1
		r.mockUser = append(r.mockUser, user)

		data,err := json.MarshalIndent(r.mockUser, ""," ")
		if err != nil{ return errs.NewUnexpectedError()}
		
		err = os.WriteFile("../../users.json",data,0644)
		if err != nil{ return errs.NewUnexpectedError()}

		return nil
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
	return model.User{}, errs.NewNotFoundError("User not found")
}
