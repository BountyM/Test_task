package models

type Product struct {
	Id         int        `bd:"id"`
	Name       string     `bd:"name"`
	Mark       int        `bd:"mark"`
	Categories []Category `bd:"categories"`
}

type Category struct {
	Id   int    `bd:"id"`
	Name string `bd:"name"`
}
