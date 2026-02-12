package service

import (
	"go-basics/internal/model"
	"go-basics/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

/* func isValidVariableName(name string) bool {
	// Validate variable name pattern
	pattern := `^[a-zA-Z_][a-zA-Z0-9_]*$`
	match, _ := regexp.MatchString(pattern, name)
	return match
} */

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(name string, age int) {

	user := model.User{
		ID: len(s.repo.GetAll()) + 1,

		Name: name,

		Age: age,
	}

	s.repo.Add(user)
}

/* func (s *UserService) CreateUser(name string, age int) {

	validName := isValidVariableName(name)

	if validName {

		user := model.User{
			ID: len(s.repo.GetAll()) + 1,

			Name: name,

			Age: age,
		}

		s.repo.Add(user)
		fmt.Println("Name is valid")
	} else {
		fmt.Println("Please provide valid name")
	}
} */

func (s *UserService) GetUSers() []model.User {
	return s.repo.GetAll()
}

func (s *UserService) DeleteUser(id int) bool {
	return s.repo.DeleteByID(id)
}

func (s *UserService) UpdateUser(id int, name string, age int) bool {
	return s.repo.UpdateByID(id, name, age)
}
