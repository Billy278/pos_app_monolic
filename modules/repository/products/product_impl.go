package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	modelsProducts "github.com/Billy278/pos_app_monolic/modules/models/products"
)

type ProductsImpl struct {
	DB *sql.DB
}

func NewProductsImpl(db *sql.DB) Product {
	return &ProductsImpl{
		DB: db,
	}
}
func (repo *ProductsImpl) RepoList(ctx context.Context, limit, offset uint64) (resProduct []modelsProducts.Products, err error) {
	fmt.Println("RepoList")
	sqlList := "SELECT id,name,stock,price,image,category_id,created_at,updated_at FROM products limit $1 offset $2"
	row, err := repo.DB.QueryContext(ctx, sqlList, limit, offset)
	if err != nil {
		return
	}
	defer row.Close()
	product := modelsProducts.Products{}
	for row.Next() {
		err = row.Scan(&product.Id, &product.Name, &product.Stock, &product.Price, &product.Image, &product.Category_id, &product.Created_At, &product.Updated_At)
		if err != nil {
			return
		}
		resProduct = append(resProduct, product)
	}
	return

}
func (repo *ProductsImpl) RepoFindByid(ctx context.Context, id uint64) (resProduct modelsProducts.Products, err error) {
	fmt.Println("RepoFindByid")
	sqlFind := "SELECT id,name,stock,price,image,category_id,created_at,updated_at FROM products WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id, &resProduct.Name, &resProduct.Stock, &resProduct.Price, &resProduct.Image, &resProduct.Category_id, &resProduct.Created_At, &resProduct.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
func (repo *ProductsImpl) RepoCreate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error) {
	fmt.Println("RepoCreate")
	sqlCreate := "INSERT INTO products(name,stock,price,image,category_id,created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
	row, err := repo.DB.QueryContext(ctx, sqlCreate, productIn.Name, productIn.Stock, productIn.Price, productIn.Image, productIn.Category_id, productIn.Created_At)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id)
		if err != nil {
			return
		}
	}

	resProduct.Name = productIn.Name
	resProduct.Stock = productIn.Stock
	resProduct.Price = productIn.Price
	resProduct.Image = productIn.Image
	resProduct.Category_id = productIn.Category_id
	resProduct.Created_At = productIn.Created_At
	return
}
func (repo *ProductsImpl) RepoUpdate(ctx context.Context, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error) {
	fmt.Println("RepoUpdate")
	sqlUpdate := "UPDATE products set name=$1,stock=$2,price=$3,image=$4,category_id=$5,updated_at=$6 WHERE id=$7"

	_, err = repo.DB.ExecContext(ctx, sqlUpdate, productIn.Name, productIn.Stock, productIn.Price, productIn.Image, productIn.Category_id, productIn.Updated_At, productIn.Id)
	if err != nil {
		return
	}
	resProduct.Id = productIn.Id
	resProduct.Name = productIn.Name
	resProduct.Stock = productIn.Stock
	resProduct.Price = productIn.Price
	resProduct.Image = productIn.Image
	resProduct.Category_id = productIn.Category_id
	resProduct.Created_At = productIn.Created_At
	return
}
func (repo *ProductsImpl) RepoDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("RepoDelete")
	sqlDelete := "DELETE FROM products WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}
	return
}

func (repo *ProductsImpl) RepoUpdateTx(ctx context.Context, db *sql.Tx, productIn modelsProducts.Products) (resProduct modelsProducts.Products, err error) {
	fmt.Println("RepoUpdateProductTX")
	sqlUpdate := "UPDATE products set name=$1,stock=$2,price=$3,image=$4,category_id=$5,updated_at=$6 WHERE id=$7"

	_, err = db.ExecContext(ctx, sqlUpdate, productIn.Name, productIn.Stock, productIn.Price, productIn.Image, productIn.Category_id, productIn.Updated_At, productIn.Id)
	if err != nil {
		return
	}
	resProduct.Id = productIn.Id
	resProduct.Name = productIn.Name
	resProduct.Stock = productIn.Stock
	resProduct.Price = productIn.Price
	resProduct.Image = productIn.Image
	resProduct.Category_id = productIn.Category_id
	resProduct.Created_At = productIn.Created_At

	return
}
func (repo *ProductsImpl) RepoFindByidTx(ctx context.Context, db *sql.Tx, id uint64) (resProduct modelsProducts.Products, err error) {
	fmt.Println("RepoFindByidProductTx")
	sqlFind := "SELECT id,name,stock,price,image,category_id,created_at,updated_at FROM products WHERE id=$1 FOR UPDATE"
	row, err := db.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resProduct.Id, &resProduct.Name, &resProduct.Stock, &resProduct.Price, &resProduct.Image, &resProduct.Category_id, &resProduct.Created_At, &resProduct.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}

	return
}
