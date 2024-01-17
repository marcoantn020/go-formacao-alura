package models

import (
	"alura/db"
	"database/sql"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	conn := db.ConnectWithDatabase()

	selectAllProducts, err := conn.Query("SELECT * FROM products ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	var products []Product

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price

		products = append(products, product)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)
	return products
}

func CreateNewProduct(name string, description string, price float64, quantity int) {
	conn := db.ConnectWithDatabase()
	insert, err := conn.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(name, description, price, quantity)
	if err != nil {
		panic(err.Error())
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)
}

func Delete(id string) {
	conn := db.ConnectWithDatabase()

	prepare, err := conn.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	_, err = prepare.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

func FindById(id string) Product {
	conn := db.ConnectWithDatabase()

	result, err := conn.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for result.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := result.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	return product
}

func Update(id int, name string, description string, price float64, quantity int) {
	conn := db.ConnectWithDatabase()
	update, err := conn.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	_, err = update.Exec(name, description, price, quantity, id)
	if err != nil {
		panic(err.Error())
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)
}
