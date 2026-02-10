package main

import (
	"bufio"
	"fmt"
	"go-basics/internal/repository"
	"go-basics/internal/service"
	"os"
	"strconv"
	"strings"
)

func main() {

	repo := repository.NewUserRepository()

	userService := service.NewUserService(repo)

	reader := bufio.NewReader(os.Stdin)


	fmt.Println("User registry CLI")
	
	fmt.Println("Create, List, Exit")


	for {
		
		input, _ :=reader.ReadString('\n')

		command := strings.TrimSpace(input)


		switch command {
		case "Create":
			fmt.Print("Skriv in ditt namn: ")

			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)


			fmt.Print("Skriv in ålder: ")

			age, _ := reader.ReadString('\n')
			age = strings.TrimSpace(age)


			ageInput, err := strconv.Atoi(age)

			if err != nil {
				fmt.Println("Age måste vara ett number")
				continue
			}
			userService.CreateUser(name,ageInput)

			fmt.Println("User skapad")

		case "List":

			users := userService.GetUSers()

			for _, u := range users {
				fmt.Println(u.ID, u.Name, u.Age)
			}

		case "Exit":
			fmt.Println("Programmet avslutas")
			return

		default:
			fmt.Println("Okänt kommando")
				
		}
	}
}