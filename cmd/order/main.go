package main

import (
	"database/sql"

	"github.com/fhlpassis/fullcycle-go-intensivo-2023-07/internal/infra/database"
	"github.com/fhlpassis/fullcycle-go-intensivo-2023-07/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
}
