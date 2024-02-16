package logging

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

var (
	client *logging.Client
	logger *logging.Logger
)

func Setup(ctx context.Context) error {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT") // Certifique-se de configurar essa variável de ambiente
	var err error

	client, err = logging.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	logName := "my-log-name" // Substitua pelo nome do seu log
	logger = client.Logger(logName)

	return nil
}

func LogInfo(ctx context.Context, msg string) {
	logger.Log(logging.Entry{Payload: msg, Severity: logging.Info})
}

func LogError(ctx context.Context, msg string, err error) {
	logger.Log(logging.Entry{Payload: msg, Severity: logging.Error})
}

func Shutdown(ctx context.Context) {
	// É importante fechar o cliente de logging para garantir que todas as entradas pendentes sejam enviadas.
	err := client.Close()
	if err != nil {
		log.Printf("Failed to close logging client: %v", err)
	}
}

// Logger retorna uma instância do logger configurado.
func Logger(ctx context.Context) *logging.Logger {
	return logger
}
