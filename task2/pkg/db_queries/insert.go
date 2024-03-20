package dbqueries

import (
	"database/sql"
	"fmt"
	"math/rand"
	"task2/models"
	"time"
)

// Заполняет бд рандомными значениями
func InsertRandom(db *sql.DB, n int) error {
	categorys := make([]models.Category, 0, n)
	for i := 0; i < n; i++ {
		temp := models.Category{Id: i, Name: "namecat_" + randSeq(4)}
		categorys = append(categorys, temp)
	}

	products := make([]models.Product, 0, n*2)
	for i := 0; i < n*2; i++ {
		temp_c := make([]models.Category, 0)
		for i := 0; i < rand.Intn(2)+2; i++ {

			temp_c = append(temp_c, categorys[rand.Intn(n)])
		}
		temp_p := models.Product{Id: i, Name: "nameproduct_" + randSeq(4), Mark: rand.Intn(10) + 1, Categories: temp_c}

		products = append(products, temp_p)
	}

	for _, p := range products {
		err := InsertProduct(db, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// возвращает рандомную строку длины n
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// добавляет продукт и категории в таблицы
func InsertProduct(db *sql.DB, p models.Product) error {
	if flag, err := IsProductValid(db, p); err != nil {
		return err
	} else if flag {
		return fmt.Errorf("product %s already exists", p.Name)
	}

	stmt1, err := db.Prepare("INSERT INTO products(id, name, mark) values($1,$2,$3)")
	if err != nil {
		return err
	}

	stmt2, err := db.Prepare("INSERT INTO categories(id, name) values($1,$2)")
	if err != nil {
		return err
	}

	stmt3, err := db.Prepare("INSERT INTO products_categories(id_product, id_categories) values($1,$2)")
	if err != nil {
		return err
	}

	stmt1.Exec(p.Id, p.Name, p.Mark)
	for _, v := range p.Categories {
		stmt2.Exec(v.Id, v.Name)
		stmt3.Exec(p.Id, v.Id)
	}

	return nil
}

// проверка на наличие продукта
func IsProductValid(db *sql.DB, p models.Product) (bool, error) {

	rows, err := db.Query("SELECT * FROM products WHERE id = $1 \n", p.Id)
	if err != nil {
		return false, err
	}

	temp := models.Product{}
	var c1 int
	var c2 string
	var c3 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3)
		if err != nil {
			return false, err
		}
		temp = models.Product{Id: c1, Name: c2, Mark: c3, Categories: []models.Category{}}
	}

	return p.Name == temp.Name, nil
}
