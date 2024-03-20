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
		log.Fatal(err)
	}

	// заполняем бд рандомными значениями
	if err := dbqueries.InsertRandom(db, 10); err != nil {
		log.Fatal(err)
	}

	// получаем значение по id
	p, err := dbqueries.GetProduct(db, 3)
	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Printf("print product id = %d\n", p.Id)
	fmt.Printf("Product: %s\nMark: %v\nCategories: ", p.Name, p.Mark)
	for _, v := range p.Categories {
		fmt.Print(v.Name + " ")
	}
	fmt.Println()
}
