package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/lucasantoniooficial/hexagonal-go/adapters/db"
	"github.com/lucasantoniooficial/hexagonal-go/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	stmt, _ := db.Prepare(`CREATE TABLE products (id TEXT, name TEXT, price REAL, status TEXT);`)
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?);`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec("1", "Product1", 10.0, "available")
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "available", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 20.0

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "disabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
