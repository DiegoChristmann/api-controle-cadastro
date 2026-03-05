-- Criar banco de dados se não existir
CREATE DATABASE controle_cadastro;

-- Conectar ao banco
\c controle_cadastro

-- Criar tabela de usuários
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Inserir dados de exemplo
INSERT INTO "user" (user_name, email) VALUES 
    ('Diego Christmann', 'diego@gmail.com'),
    ('Maria Santos', 'maria@example.com'),
    ('Pedro Oliveira', 'pedro@example.com'),
    ('Ana Costa', 'ana@example.com'),
    ('Carlos Souza', 'carlos@example.com')
ON CONFLICT DO NOTHING;
