package db

type User struct {
	ID       int
	UserName string
	Email    string
}

var MockUsers = []User{
	{ID: 1, UserName: "Diego Christmann", Email: "diego@gmail.com"},
	{ID: 2, UserName: "Maria Santos", Email: "maria@example.com"},
	{ID: 3, UserName: "Pedro Oliveira", Email: "pedro@example.com"},
	{ID: 4, UserName: "Ana Costa", Email: "ana@example.com"},
	{ID: 5, UserName: "Carlos Souza", Email: "carlos@example.com"},
}

// GetAllUsers retorna todos os usuários mock
func GetAllUsers() []User {
	return MockUsers
}

// GetUserByID retorna um usuário específico pelo ID
func GetUserByID(id int) *User {
	for i := range MockUsers {
		if MockUsers[i].ID == id {
			return &MockUsers[i]
		}
	}
	return nil
}

// GetUserByEmail retorna um usuário específico pelo email
func GetUserByEmail(email string) *User {
	for i := range MockUsers {
		if MockUsers[i].Email == email {
			return &MockUsers[i]
		}
	}
	return nil
}
