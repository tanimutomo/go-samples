package database

import "github.com/tanimutomo/gin-api-server-ja-tutorials/clean-architecture/src/app/domain"

type UserRespository struct {
	SqlHandler
}

func (repo *UserRespository) Store(u domain.User) (id int, err error) {
	result, err := repo.Executre(
		"INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRespository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return
}

func (repo *UserRespository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var fistName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}
