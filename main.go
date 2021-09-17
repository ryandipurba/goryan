package main

import (
	"baju/repository"
	usecase "baju/use_case"
	"fmt"
)

func main() {

	r := repository.NewBajuRepository()
	us := usecase.NewBajuUseCase(r)

	us.Addbaju(1, "Bilabong", 128000)
	us.Addbaju(2, "Adidas", 223000)
	us.Addbaju(3, "Nike", 124000)
	us.Addbaju(4, "Balenciaga", 133000)
	us.Addbaju(5, "Hugo", 523000)
	us.Addbaju(6, "Boss", 143000)
	us.Addbaju(7, "Blueberry", 123000)
	us.Addbaju(8, "Crocodile", 143000)
	us.Addbaju(9, "Lacoste", 125000)
	us.Addbaju(10, "Lv", 133000)

	us.DeleteBaju(5)
	// fmt.Printf("List Baju  %v", us.BajuLimit(5))
	fmt.Printf("List Baju  %v", us.GetList())
}
