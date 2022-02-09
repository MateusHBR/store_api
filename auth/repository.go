package auth

import (
	"database/sql"

	"github.com/MateusHBR/mallus_api/server"
	"github.com/google/uuid"
)

type repository interface {
	createAccount(CreateAccount) (CreateAccount, error)
}

type authRepository struct {
	*sql.DB
}

func (r authRepository) createAccount(data CreateAccount) (CreateAccount, error) {
	data.ID = uuid.New().String()
	createdAt := server.TimeNow()
	updatedAt := createdAt

	sqlStmt := `
	INSERT INTO users(id, first_name, last_name, email, password, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.Exec(sqlStmt, data.ID, data.FirstName, data.LastName, data.Email, data.Password, createdAt, updatedAt)

	if err != nil {
		return CreateAccount{}, err
	}

	return CreateAccount{
		ID:        data.ID,
		Email:     data.Email,
		Password:  data.Password,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}, nil
}
