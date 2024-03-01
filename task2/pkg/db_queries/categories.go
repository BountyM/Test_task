package dbqueries

import (
	"database/sql"
	"log"
	"task2/models"
)

// возвращает категорию по id
func GetCategory(db *sql.DB, id_category int) (models.Category, error) {

	rows, err := db.Query("SELECT * FROM categories where id = $1\n", id_category)
	if err != nil {
		log.Println("Query:", err)
		return models.Category{}, err
	}
	defer rows.Close()

	var temp models.Category
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Name)
		if err != nil {
			log.Println(err)
			return models.Category{}, err
		}
	}
	return temp, nil
}

// возвращает map[category.id]category.name
func GetMapCatIdsOfName(db *sql.DB) (map[int]string, error) {
	catIdsOfName := make(map[int]string)

	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		log.Println("Query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i int
		var n string
		err = rows.Scan(&i, &n)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		catIdsOfName[i] = n
	}

	return catIdsOfName, nil
}
