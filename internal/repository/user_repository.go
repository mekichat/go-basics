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

func (r *UserRepository) DeleteByID(id int) bool {

	for i, user := range r.users {

		if user.ID == id {

			r.users = append(
				r.users[:i],
				r.users[i+1:]...,
			)
			
			return true
		}
	}

	return false
	
}


func (r *UserRepository) UpdateByID(id int, newName string, newAge int) bool {
	

	for i := range r.users {


		if r.users[i].ID == id {

			r.users[i].Name = newName

			r.users[i].Age = newAge

			return true

		}

	}

	return false

}