DROP TABLE IF EXISTS products_categories, products, categories;
DROP VIEW IF EXISTS v;

CREATE DATABASE dbname;

CREATE TABLE products (
id INT PRIMARY KEY,
name VARCHAR NOT NULL,
mark INT
);

CREATE TABLE categories (
id INT PRIMARY KEY,
name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE products_categories (
id_product INT,
FOREIGN KEY (id_product)  REFERENCES products (id),
id_categories INT, 
FOREIGN KEY (id_categories)  REFERENCES categories (id)
);

ALTER TABLE products_categories ADD CONSTRAINT pc_feature_unique  UNIQUE (id_product, id_categories);

CREATE OR REPLACE VIEW v AS 
	(select p.name as pn,  p.mark, c.name as cn
		from categories c
		left join products_categories pc 
		on pc.id_categories = c.id 
		left join products p 
		on p.id = pc.id_product);
