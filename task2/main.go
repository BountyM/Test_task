package main

import (
	"fmt"
	"log"
	"task2/models"
	dbqueries "task2/pkg/db_queries"

	_ "github.com/lib/pq"
)

func main() {
	// коннуктимся к бд
	db, err := dbqueries.ConnectPostgres()
	if err != nil {
		log.Println(err)
		return
	}

	// заполняем бд рандомными значениями
	dbqueries.InsertRandom(db, 10)

	// получаем значение по id
	p, err := dbqueries.GetProduct(db, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("print product id = 3")
	PtintProduct(p)

	fmt.Println("-------------")
	fmt.Println("min")
	mapa, err := dbqueries.MinMaxMarkInCats(db, "min")
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range mapa {
		fmt.Println(k, mapa[k])
	}

	fmt.Println("-------------")
	fmt.Println("max")
	mapa, err = dbqueries.MinMaxMarkInCats(db, "max")
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range mapa {
		fmt.Println(k, mapa[k])
	}
}

func PtintProduct(p models.Product) {
	fmt.Printf("Product: %s\nMark: %v\nCategories: ", p.Name, p.Mark)
	for _, v := range p.Categories {
		fmt.Print(v.Name + " ")
	}
	fmt.Println()
}
