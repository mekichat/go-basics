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
	
	fmt.Println("Create, List ,Delete ,Update , Exit")


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


		case "Delete":
			fmt.Print("Skriv ID på användaren du vill radera: ")
		
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)


			idInput, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("ID måste vara ett numbber")
				continue
			}

			deleted := userService.DeleteUser(idInput)

			if deleted {
				fmt.Println("Användaren raderad")
			} else {
					fmt.Println("Användaren hittades inte")	
			}


		case "Update":

			fmt.Print("Skriv ID på användaren du vill uppdatera: ")
		
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)


			idInput, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("ID måste vara ett numbber")
				continue
			}


			fmt.Print("Type in new name: ")

			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)


			fmt.Print("Type in new age: ")

			newAge, _ := reader.ReadString('\n')
			newAge = strings.TrimSpace(newAge)


			ageInput, err := strconv.Atoi(newAge)

			if err != nil {
				fmt.Println("ID måste vara ett numbber")
				continue
			}

			updated := userService.UpdateUser(idInput, newName, ageInput)

			if updated {
				fmt.Println("User uppdaterad")
			} else {
				fmt.Println("Användaren kunde inte hittas..")
			}

		case "Exit":
			fmt.Println("Programmet avslutas")
			return

		default:
			fmt.Println("Okänt kommando")
				
		}
	}
}