package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/thesunonthesky/GoBackendPackage/model"
)

type UserRepository interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
}

type UserRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Create(user model.User) error {
	query := `Insert Into users(name,email) values(?,?)`
	_, error := r.db.Exec(query, user.Name, user.Email)


	return error
}


func (r *UserRepositoryDB) GetAll() ([]model.User, error){
	users := []model.User{}
	query := `Select * from users`
	err  := r.db.Select(&users,query)
	return users,err
}
