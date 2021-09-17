package main

import (
	"baju/repository"
	usecase "baju/use_case"
	"fmt"
)

func main() {

	r := repository.NewBajuRepository()
	us := usecase.NewBajuUseCase(r)

	// us.AddBaju(1, "asd")
	fmt.Printf("List Baju  %v", us.GetList())
}
