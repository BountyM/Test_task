package dbqueries

import (
	"database/sql"
	"log"
	"task2/models"
)

// получает продукт из бд по id
func GetProduct(db *sql.DB, id int) (models.Product, error) {

	rows, err := db.Query(`
	SELECT p.id, p.name, p.mark, c.id , c.name 
	FROM  products_categories pc
	LEFT JOIN products p
	ON p.id =pc.id_product 
	left join categories c 
	on pc.id_categories =c.id 
	where p.id = $1
	order by p.id;
	`, id)

	if err != nil {
		log.Println("Query:", err)
		return models.Product{}, err
	}
	defer rows.Close()

	temp_product := models.Product{}
	temp_categories := models.Category{}
	cat_slice := make([]models.Category, 0, 10)

	for rows.Next() {
		err = rows.Scan(&temp_product.Id, &temp_product.Name, &temp_product.Mark, &temp_categories.Id, &temp_categories.Name)
		if err != nil {
			log.Println(err)
			return models.Product{}, err
		}
		cat_slice = append(cat_slice, temp_categories)
	}
	temp_product.Categories = cat_slice
	return temp_product, nil

}
