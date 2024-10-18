package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// String JSON com caracteres de escape
	jsonString := `{"name":"John","age":30,"city":"New York"}`

	// Mapa genérico para armazenar o JSON
	var result map[string]interface{}

	// Parsear a string JSON
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("Erro ao parsear o JSON:", err)
		return
	}

	// Serializar novamente sem os caracteres de escape e formatado
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar o JSON:", err)
		return
	}

	// Criar arquivo físico .json
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Escrever o JSON formatado no arquivo
	_, err = file.Write(prettyJSON)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo JSON salvo com sucesso!")
}

// Função para fazer upload do arquivo para o S3
func uploadToS3(fileName string) error {
	// Configuração da sessão AWS para o LocalStack
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:         aws.String("http://localhost:4566"), // Endpoint do LocalStack
		S3ForcePathStyle: aws.Bool(true),                      // Necessário para LocalStack
	})
	if err != nil {
		return fmt.Errorf("erro ao criar sessão AWS: %w", err)
	}

	// Criar serviço S3
	s3Client := s3.New(sess)

	// Abrir o arquivo que será enviado
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo para upload: %w", err)
	}
	defer file.Close()

	// Obter informações do arquivo
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("erro ao obter informações do arquivo: %w", err)
	}

	// Preparar a solicitação de upload
	bucket := "meu-bucket" // Nome do bucket
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(fileName), // Nome do arquivo no S3
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String("application/json"),
	})
	if err != nil {
		return fmt.Errorf("erro ao fazer upload para o S3: %w", err)
	}

	return nil
}
