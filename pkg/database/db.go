package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func SetupDatabase() error {
	var err error
	databaseURL := os.Getenv("DATABASE_URL") // Certifique-se de que DATABASE_URL está configurado nas variáveis de ambiente

	DB, err = pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	// Testa a conexão
	if err = DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return nil
}

func FetchCNPJs() ([]string, error) {
	rows, err := DB.Query(context.Background(), "SELECT cnpj FROM configuration")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var cnpjs []string
	for rows.Next() {
		var cnpj string
		if err := rows.Scan(&cnpj); err != nil {
			return nil, fmt.Errorf("rows.Scan failed: %v", err)
		}
		cnpjs = append(cnpjs, cnpj)
	}
	return cnpjs, rows.Err()
}

// Lembre-se de fechar a conexão com o banco de dados quando o programa terminar
func CloseDatabase() {
	DB.Close()
}

// InsertData insere os dados recebidos da API para um CNPJ específico no banco de dados.
func InsertData(ctx context.Context, cnpj string, data string) error {
	commandTag, err := DB.Exec(ctx, "INSERT INTO pharma_data (cnpj, data) VALUES ($1, $2)", cnpj, data)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no rows were inserted")
	}
	return nil
}
