package service

import (
	"go-basics/internal/model"
	"go-basics/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService{
	return &UserService{
		repo: r,
	}
}




func (s *UserService) CreateUser(name string, age int) {
	user := model.User {
		ID: len(s.repo.GetAll()) + 1,

		Name: name,

		Age: age,
	}

	s.repo.Add(user)
}


func (s *UserService) GetUSers() []model.User {
	return s.repo.GetAll()
}
