package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pauloscardoso/go-intensivo-jul/internal/infra/database"
	"github.com/pauloscardoso/go-intensivo-jul/internal/usecase"
)

type Car struct {
	Model string
	Color string
}

// método
// func (c Car) Start() {
// 	println(c.Model + " is moving")
// }

// função
// func soma(x, y int) int {
// 	return x + y
// }

func (c *Car) ChangeColor(color string) { // usar * para acessar o valor original na memória e não a cópia
	c.Color = color // cópia do color original
	println("New Color: " + c.Color)
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	or := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(or)

	input := usecase.OrderInput{
		ID:    "12345",
		PRICE: 100,
		TAX:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	// uc.Execute(input)

	// car := Car{ // := declara e atribui
	// 	Model: "Audi",
	// 	Color: "Red",
	// }
	// car.Start()
	// car.ChangeColor("Blue")
	// println(car.Color)
	// car.Model = "Fiat"
	// println(car.Model)

	// order, err := entity.NewOrder("123", 100, 10)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }
}
