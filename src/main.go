package main

import (
	
	"fmt"

	"github.com/Tiago-Salles/go-user-service/src/domain/use_cases"

)


func main(){
	fmt.Println("APP IS RUNNING")

	useCase := use_cases.UserUseCase{}
	user := useCase.Register()
	fmt.Println(user.Name)
	fmt.Println(user.Email)
}