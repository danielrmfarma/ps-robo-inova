package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/danielrmfarma/ps-robo-inova/pkg/api"
	"github.com/danielrmfarma/ps-robo-inova/pkg/database"
	"github.com/danielrmfarma/ps-robo-inova/pkg/logging"
)

func main() {
	ctx := context.Background()

	// Configuração de logging
	if err := logging.Setup(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set up logging: %v\n", err)
		os.Exit(1)
	}
	defer logging.Shutdown(ctx)

	// Inicialização do banco de dados
	if err := database.SetupDatabase(); err != nil {
		logging.LogError(ctx, "Failed to set up database: %v", err)
		os.Exit(1)
	}
	defer database.CloseDatabase()

	// Aqui você chamaria suas funções de inicialização e execução
	// Por exemplo: buscar CNPJs e disparar as operações
	cnpjs, err := database.FetchCNPJs()
	if err != nil {
		logging.LogError(ctx, "Failed to fetch CNPJs: %v", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, cnpj := range cnpjs {
		wg.Add(1)
		go func(cnpj string) {
			defer wg.Done()
			api.FetchApiDataAndStore(ctx, cnpj) // Execute a operação de fetch e store
		}(cnpj)
	}

	wg.Wait()
	logging.LogInfo(ctx, "All operations completed successfully.")
}
