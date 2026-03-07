package repository

import (
	"api-controle-cadastro/model"
	"database/sql"
	"fmt"
)

type userRepository struct { // Renamed from UserRepository to avoid conflict with the interface
	connection *sql.DB
}

// ...existing code...

type UserRepository interface { // Keep this as the interface
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (int, error)
	GetUserById(id_user int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	DeleteUser(id_user int) error
	UpdateUser(user model.User) error
}

// ...existing code...

func NewUserRepository(connection *sql.DB) UserRepository { // Now returns the interface type
	return &userRepository{ // Return a pointer to the struct, which implements the interface
		connection: connection,
	}
}

func (pr *userRepository) GetUsers() ([]model.User, error) {
	rows, err := pr.connection.Query("SELECT id, user_name, email FROM \"user\" ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %w", err)
	}
	defer rows.Close()

	var userList []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear usuário: %w", err)
		}
		userList = append(userList, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre usuários: %w", err)
	}

	return userList, nil
}

func (pr *userRepository) CreateUser(user model.User) (int, error) {
	var id int
	err := pr.connection.QueryRow(
		"INSERT INTO \"user\" (user_name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return id, nil
}

func (pr *userRepository) GetUserById(id_user int) (*model.User, error) {
	var user model.User
	err := pr.connection.QueryRow(
		"SELECT id, user_name, email FROM \"user\" WHERE id = $1",
		id_user,
	).Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return &user, nil
}

func (pr *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := pr.connection.QueryRow(
		"SELECT id, user_name, email FROM \"user\" WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return &user, nil
}

func (pr *userRepository) DeleteUser(id_user int) error {
	result, err := pr.connection.Exec("DELETE FROM \"user\" WHERE id = $1", id_user)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado", id_user)
	}

	return nil
}

func (pr *userRepository) UpdateUser(user model.User) error {
	// Monta query dinâmica apenas com os campos enviados
	query := "UPDATE \"user\" SET "
	args := []interface{}{}
	argPos := 1

	if user.Name != "" {
		query += fmt.Sprintf("user_name = $%d", argPos)
		args = append(args, user.Name)
		argPos++
	}

	if user.Email != "" {
		if len(args) > 0 {
			query += ", "
		}
		query += fmt.Sprintf("email = $%d", argPos)
		args = append(args, user.Email)
		argPos++
	}

	// Se nenhum campo foi enviado, não há o que atualizar
	if len(args) == 0 {
		return fmt.Errorf("nenhum campo para atualizar")
	}

	query += fmt.Sprintf(" WHERE id = $%d", argPos)
	args = append(args, user.ID)

	result, err := pr.connection.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado", user.ID)
	}

	return nil
}
