package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ConnectDB cria e retorna uma conexão com o banco de dados PostgreSQL
func ConnectDB() (*sql.DB, error) {
	// Usa variáveis de ambiente ou valores padrão
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "dadosTestes")
	sslmode := getEnv("DB_SSLMODE", "disable")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	// Configurar pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	fmt.Printf("Conectado ao banco de dados PostgreSQL: %s@%s:%s/%s\n", user, host, port, dbname)

	// Criar tabela se não existir
	err = Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("erro ao migrar banco de dados: %w", err)
	}

	return db, nil
}

// Migrate cria as tabelas necessárias no banco de dados
func Migrate(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS "user" (
			id SERIAL PRIMARY KEY,
			user_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela user: %w", err)
	}

	fmt.Println("Tabela 'user' verificada/criada com sucesso")
	return nil
}
