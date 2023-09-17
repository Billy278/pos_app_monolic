package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	modelsCategories "github.com/Billy278/pos_app_monolic/modules/models/categories"
)

type CategoriesImpl struct {
	DB *sql.DB
}

func NewCategoriesImpl(db *sql.DB) Categories {
	return &CategoriesImpl{
		DB: db,
	}
}

func (repo *CategoriesImpl) RepoList(ctx context.Context, limit, offset uint64) (resCategories []modelsCategories.Categories, err error) {
	fmt.Println("RepoList")
	sqlList := "SELECT id,name,created_at,updated_at FROM categories Limit $1 offset $2"
	row, err := repo.DB.QueryContext(ctx, sqlList, limit, offset)
	if err != nil {
		return
	}
	defer row.Close()
	categories := modelsCategories.Categories{}
	for row.Next() {
		err = row.Scan(&categories.Id, &categories.Name, &categories.Created_At, &categories.Updated_At)
		if err != nil {
			return
		}
		resCategories = append(resCategories, categories)
	}

	return
}
func (repo *CategoriesImpl) RepoFindByid(ctx context.Context, id uint64) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("RepoFindByid")
	sqlFind := "SELECT id,name,created_at,updated_at FROM categories WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resCategories.Id, &resCategories.Name, &resCategories.Created_At, &resCategories.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")

	}
	return
}
func (repo *CategoriesImpl) RepoCreate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("RepoCreate")
	sqlCreate := "INSERT INTO categories(name,created_at) VALUES ($1,$2)"
	_, err = repo.DB.ExecContext(ctx, sqlCreate, categoriesIn.Name, categoriesIn.Created_At)
	if err != nil {
		return
	}
	resCategories.Name = categoriesIn.Name

	return
}
func (repo *CategoriesImpl) RepoUpdate(ctx context.Context, categoriesIn modelsCategories.Categories) (resCategories modelsCategories.Categories, err error) {
	fmt.Println("RepoUpdate")
	sqlUpdate := "UPDATE categories set name=$1, updated_at=$2 RETURNING id"
	row, err := repo.DB.QueryContext(ctx, sqlUpdate, categoriesIn.Name, categoriesIn.Updated_At)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&resCategories.Id)
		if err != nil {
			return
		}
		resCategories.Name = categoriesIn.Name
	}
	return
}
func (repo *CategoriesImpl) RepoDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("RepoDelete")
	sqlDelete := "DELETE FROM categories WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}
	return
}
