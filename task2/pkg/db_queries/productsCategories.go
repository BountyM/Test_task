package dbqueries

import (
	"database/sql"
	"fmt"
	"log"
)

// возвращает [categories.id]
func GetCatIdsFromPC(db *sql.DB, id_product int) ([]int, error) {

	rows, err := db.Query("SELECT * FROM products_categories where id_product = $1\n", id_product)
	if err != nil {
		log.Println("Query:", err)
		return nil, err
	}
	defer rows.Close()

	cat_ids := make([]int, 0)
	for rows.Next() {
		var temp [2]int
		err = rows.Scan(&temp[0], &temp[1])
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cat_ids = append(cat_ids, temp[1])
	}
	return cat_ids, nil
}

// возвращает map[category.name]product.name с максимальным mark
func MaxMarkInCats(db *sql.DB) (map[string]string, error) {

	rows, err := db.Query(`
	select cn, pn, mark from 
	(select p.name as pn,  p.mark, c.name as cn
	from categories c
	left join products_categories pc 
	on pc.id_categories = c.id 
	left join products p 
	on p.id = pc.id_product ) t1
	WHERE mark = ( SELECT MAX( t2.mark )
				FROM 
				(select p.name as pn,  p.mark, c.name as cn
				from categories c
				left join products_categories pc 
				on pc.id_categories = c.id 
				left join products p 
				on p.id = pc.id_product ) t2
				WHERE t1.cn = t2.cn )
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
	select cn, pn, mark from v as t1
		WHERE mark = ( SELECT MIN( t2.mark )
					FROM 
					v as t2
					WHERE t1.cn = t2.cn )
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

// возвращает map[category.id][]product.ids
func GetMapCatsOfProductIds(db *sql.DB) (map[int][]int, error) {
	mapCatsOfProductIds := make(map[int][]int)

	rows, err := db.Query("SELECT * FROM products_categories")
	if err != nil {
		log.Println("Query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c, p int
		err = rows.Scan(&p, &c)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		mapCatsOfProductIds[c] = append(mapCatsOfProductIds[c], p)
	}

	return mapCatsOfProductIds, nil
}
