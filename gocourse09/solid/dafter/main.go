// The program follows the Dependency Inversion Principle from SOLID.

package main

import (
	"errors"
	"fmt"
)

type MySQL struct{}

func (db MySQL) Query() (any, error) {
	return []string{"alex", "john", "mike"}, nil
}

type PostgreSQL struct{}

func (db PostgreSQL) Query() (any, error) {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}, nil
}

type DBConn interface {
	Query() (any, error)
}

type UsersRepository struct {
	db DBConn
}

func (r UsersRepository) Users() ([]string, error) {
	var users []string
	res, err := r.db.Query()
	if err != nil {
		return nil, err
	}

	switch coll := res.(type) {
	case map[string]string:
		for _, u := range coll {
			users = append(users, u)
		}
		return users, nil
	case []string:
		return coll, nil
	}

	return nil, errors.New("failed to retrieve users")
}

func main() {
	var mysqlDB MySQL
	repoMySQL := UsersRepository{db: mysqlDB}

	var postgreSQLDB PostgreSQL
	repoPostgreSQL := UsersRepository{db: postgreSQLDB}

	repositories := []UsersRepository{repoMySQL, repoPostgreSQL}
	for _, repo := range repositories {
		users, err := repo.Users()
		if err != nil {
			fmt.Println("Failed to retrieve users from DB:", err)
			return
		}
		fmt.Println("Users from DB:", users)
	}
}
