package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	modelsUser "github.com/Billy278/pos_app_monolic/modules/models/users"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}
func (repo *UserRepoImpl) RepoList(ctx context.Context, limit, offset uint64) (resUser []modelsUser.User, err error) {
	fmt.Println("RepoList")
	sqlList := "SELECT id,name,role,email,created_at,updated_at FROM users Limit $1 offset $2"
	row, err := repo.DB.QueryContext(ctx, sqlList, limit, offset)
	if err != nil {
		return
	}
	defer row.Close()
	user := modelsUser.User{}
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Created_At, &user.Updated_At)
		if err != nil {
			return
		}
		resUser = append(resUser, user)
	}
	return
}
func (repo *UserRepoImpl) RepoFindByid(ctx context.Context, id uint64) (resUser modelsUser.User, err error) {
	fmt.Println("RepoFindByid")
	sqlFind := "SELECT id,name,role,email,created_at,updated_at FROM users WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, id)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&resUser.Id, &resUser.Name, &resUser.Role, &resUser.Email, &resUser.Created_At, &resUser.Updated_At)
		if err != nil {
			return
		}
	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
func (repo *UserRepoImpl) RepoCreate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error) {
	fmt.Println("RepoCreate")
	sqlCreate := "INSERT INTO users(name,role,email,username,password,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err = repo.DB.ExecContext(ctx, sqlCreate, userIn.Name, userIn.Role, userIn.Email, userIn.Username, userIn.Password, userIn.Created_At, userIn.Updated_At)
	if err != nil {
		return
	}
	resUser.Name = userIn.Name
	resUser.Username = userIn.Username
	resUser.Password = userIn.Password
	return
}
func (repo *UserRepoImpl) RepoUpdate(ctx context.Context, userIn modelsUser.User) (resUser modelsUser.User, err error) {
	fmt.Println("RepoUpdate")
	sqlUpdate := "UPDATE users SET name=$1,role=$2,email=$3,updated_at=$4 WHERE id=$5"
	_, err = repo.DB.ExecContext(ctx, sqlUpdate, userIn.Name, userIn.Role, userIn.Email, userIn.Updated_At, userIn.Id)
	if err != nil {
		return
	}
	resUser.Name = userIn.Name
	resUser.Email = userIn.Email
	resUser.Role = userIn.Role
	resUser.Updated_At = userIn.Updated_At

	return
}
func (repo *UserRepoImpl) RepoDelete(ctx context.Context, id uint64) (err error) {
	fmt.Println("RepoDelete")
	sqlDelete := "DELETE FROM users WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return
	}

	return
}

func (repo *UserRepoImpl) RepoFindUser(ctx context.Context, username string) (err error) {
	fmt.Println("RepoFindUser")
	sqlFind := "SELECT username,password FROM users WHERE username=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, username)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = errors.New("NOT FOUND")
	}
	return
}

func (repo *UserRepoImpl) RepoFindUsernameToLogin(ctx context.Context, username string) (resUser modelsUser.User, err error) {
	fmt.Println("RepoFindUsernameToLogin")
	sqlFind := "SELECT id,name,role,email,username,password FROM users WHERE username=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, username)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&resUser.Id, &resUser.Name, &resUser.Role, &resUser.Email, &resUser.Username, &resUser.Password)
		if err != nil {
			return
		}

	} else {
		err = errors.New("NOT FOUND")
	}
	return
}
