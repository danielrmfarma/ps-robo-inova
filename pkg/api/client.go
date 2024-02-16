package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/danielrmfarma/ps-robo-inova/pkg/database"
	"github.com/danielrmfarma/ps-robo-inova/pkg/logging"
)

// FetchApiDataAndStore faz a chamada para a API externa e armazena a resposta no banco de dados.
func FetchApiDataAndStore(ctx context.Context, cnpj string) {
	url := fmt.Sprintf("http://api.example.com/data?cnpj=%s", cnpj) // Substitua pela URL real
	resp, err := http.Get(url)
	if err != nil {
		logging.LogError(ctx, fmt.Sprintf("Failed to fetch data for CNPJ %s", cnpj), err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.LogError(ctx, fmt.Sprintf("Failed to read response body for CNPJ %s", cnpj), err)
		return
	}

	// Processamento adicional dos dados pode ser feito aqui, como transformar os dados JSON em um formato específico
	data := string(body) // Exemplo simples, considerando que a resposta é texto puro

	// Substitua pelo seu pacote e função de inserção real
	if err = database.InsertData(ctx, cnpj, data); err != nil {
		logging.LogError(ctx, fmt.Sprintf("Failed to insert data for CNPJ %s", cnpj), err)
		return
	}

	logging.LogInfo(ctx, fmt.Sprintf("Data inserted successfully for CNPJ %s", cnpj))
}
