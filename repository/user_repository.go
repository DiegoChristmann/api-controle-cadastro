package repository

import (
	"api-controle-cadastro/db"
	"api-controle-cadastro/model"
	"database/sql" // Assuming this is the connection type; adjust if needed
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
	DeleteUser(id_user int) error
}

// ...existing code...

func NewUserRepository(connection *sql.DB) UserRepository { // Now returns the interface type
	return &userRepository{ // Return a pointer to the struct, which implements the interface
		connection: connection,
	}
}

func (pr *userRepository) GetUsers() ([]model.User, error) { // Updated receiver
	mockUsers := db.GetAllUsers()
	var userList []model.User

	for _, u := range mockUsers {
		userList = append(userList, model.User{
			ID:    u.ID,
			Name:  u.UserName,
			Email: u.Email,
		})
	}

	return userList, nil
}

func (pr *userRepository) CreateUser(user model.User) (int, error) { // Updated receiver
	var id int
	query, err := pr.connection.Prepare("INSERT INTO user " +
		"(user_name, email) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Email).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *userRepository) GetUserById(id_user int) (*model.User, error) { // Updated receiver
	mockUser := db.GetUserByID(id_user)
	if mockUser == nil {
		return nil, nil
	}

	return &model.User{
		ID:    mockUser.ID,
		Name:  mockUser.UserName,
		Email: mockUser.Email,
	}, nil
}

func (pr *userRepository) DeleteUser(id_user int) error { // Updated receiver; properly implemented
	query, err := pr.connection.Prepare("DELETE FROM user WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(id_user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado", id_user)
	}

	// Optional: Also remove from mock if using dual logic (comment out if not needed)
	for i, u := range db.MockUsers {
		if u.ID == id_user {
			db.MockUsers = append(db.MockUsers[:i], db.MockUsers[i+1:]...)
			break
		}
	}

	return nil
}

// Remove any old misplaced code (e.g., the query.Exec and for loop that were after GetUserById)
