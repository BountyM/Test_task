package dbqueries

import (
	"database/sql"
	"fmt"
	"log"
)

// возвращает map[category.name]product.name с максимальным mark
func MaxMarkInCats(db *sql.DB) (map[string]string, error) {

	rows, err := db.Query(`
	WITH cte AS (select p.id , p.name , p.mark, c.id as id_cat, c.name as name_cat,
		RANK() OVER ( PARTITION BY c.id
					ORDER BY mark DESC
					) AS r
		from products p 
		left join products_categories pc on p.id =pc.id_product 
		left join categories c on pc.id_categories =c.id 
		)
		SELECT name_cat, name, mark
		FROM cte
		WHERE r = 1
	`)
	if err != nil {
		log.Println("Query:", err)
		return nil, err
	}

	temp := make(map[string]string)

	for rows.Next() {
		var cn, pn string
		var t int
		err = rows.Scan(&cn, &pn, &t)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		temp[cn] = pn
	}

	return temp, nil
}

// возвращает map[category.name]product.name с минимальным mark
func MinMarkInCats(db *sql.DB) (map[string]string, error) {

	rows, err := db.Query(`
	WITH cte AS (select p.id , p.name , p.mark, c.id as id_cat, c.name as name_cat,
		RANK() OVER ( PARTITION BY c.id
					ORDER BY mark ASC
					) AS r
		from products p 
		left join products_categories pc on p.id =pc.id_product 
		left join categories c on pc.id_categories =c.id 
		)
		SELECT name_cat, name, mark
		FROM cte
		WHERE r = 1
	`)
	if err != nil {
		log.Println("Query:", err)
		return nil, err
	}

	temp := make(map[string]string)

	for rows.Next() {
		var cn, pn string
		var t int
		err = rows.Scan(&cn, &pn, &t)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		temp[cn] = pn
	}

	return temp, nil
}

// вызывает либо MinMarkInCats либо MMaxMarkInCats
func MinMaxMarkInCats(db *sql.DB, minORmax string) (map[string]string, error) {
	if minORmax == "min" {
		return MinMarkInCats(db)
	} else if minORmax == "max" {
		return MaxMarkInCats(db)
	}

	return nil, fmt.Errorf("wrong value of minOrMax")
}
