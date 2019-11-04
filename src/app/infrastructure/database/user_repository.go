package database

import "app/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int , err error) {
	result, err := repo.Execute(
        "INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
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

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
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
