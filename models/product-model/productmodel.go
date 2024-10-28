package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
	SELECT 
		p.id,
		p.name,
		c.name as category_name,
		p.stock,
		p.description,
		p.created_at,
		p.updated_at
	FROM products p
	JOIN categories c ON p.category_id = c.id`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.Created_at,
			&product.Updated_at); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	res, err := config.DB.Exec(`
		INSERT INTO products (name, category_id, stock, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		product.Created_at,
		product.Updated_at,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT
			p.id,
			p.name,
			c.name as category_name,
			p.stock,
			p.description,
			p.created_at,
			p.updated_at
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.Created_at,
		&product.Updated_at,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		UPDATE products SET 
			name = ?, 
			category_id = ?, 
			stock = ?, 
			description = ?, 
			updated_at = ? 
		WHERE id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		product.Updated_at,
		id,
	)

	if err != nil {
		panic(err)
	}

	res, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}

	return res > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}