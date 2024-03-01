package dbqueries

import (
	"database/sql"
	"log"
	"task2/models"
)

// получает продукт из бд по id
func GetProduct(db *sql.DB, id int) (models.Product, error) {
	cat_ids, err := GetCatIdsFromPC(db, id)
	if err != nil {
		return models.Product{}, err
	}

	categories := make([]models.Category, 0, len(cat_ids))

	for _, i := range cat_ids {
		temp, err := GetCategory(db, i)
		if err != nil {
			return models.Product{}, err
		}
		categories = append(categories, temp)
	}

	rows, err := db.Query("SELECT * FROM products where id = $1\n", id)
	if err != nil {
		log.Println("Query:", err)
		return models.Product{}, err
	}
	defer rows.Close()

	var temp models.Product
	temp.Categories = categories

	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Mark)
		if err != nil {
			log.Println(err)
			return models.Product{}, err
		}
	}
	return temp, nil

}

// возвращает maps map[product.id]mark и map[product.id]name
func GetMapProductIdsOfMarkAndName(db *sql.DB) (map[int]int, map[int]string, error) {
	productIdsOfMark := make(map[int]int)
	productIdsOfName := make(map[int]string)

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Println("Query:", err)
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i, m int
		var n string
		err = rows.Scan(&i, &n, &m)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
		productIdsOfMark[i] = m
		productIdsOfName[i] = n
	}

	return productIdsOfMark, productIdsOfName, nil
}
