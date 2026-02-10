package repository

import "go-basics/internal/model"

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository{
	return &UserRepository{
		users: []model.User{},
	}
}


func (r *UserRepository) Add(user model.User){
	r.users = append(r.users, user)
}

func (r *UserRepository) GetAll() []model.User {
	return r.users
}