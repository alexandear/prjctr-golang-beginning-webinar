// The program violates the Dependency Inversion Principle from SOLID.

package main

import (
	"fmt"
)

type MySQL struct{}

func (db MySQL) QueryMySQL() ([]string, error) {
	return []string{"alex", "john", "mike"}, nil
}

type PostgreSQL struct{}

func (db PostgreSQL) QueryPostgreSQL() (map[string]string, error) {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}, nil
}

type UsersRepository struct {
	db MySQL
	// db PostgreSQL
}

func (r UsersRepository) Users() ([]string, error) {
	res, err := r.db.QueryMySQL() // res := r.db.QueryPostgreSQL()
	if err != nil {
		return nil, err
	}

	var users []string
	for _, u := range res {
		users = append(users, u)
	}
	return users, nil
}

func main() {
	mysqlDB := MySQL{}
	// postgreSQLDB := PostgreSQL{}

	repo := UsersRepository{db: mysqlDB}
	// repo := UsersRepository{db: postgreSQLDB}

	users, err := repo.Users()
	if err != nil {
		fmt.Println("Failed to retrieve users from DB:", err)
		return
	}
	fmt.Println("Users from PostgreSQL DB:", users)

	// fmt.Println("Users from MySQL DB:", users)
}
