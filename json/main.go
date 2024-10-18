package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// String JSON com barras invertidas antes dos colchetes e quebras de linha
	escapedJSONString := `[\n {\n \"name\": \"Carlos\", \n \"age\": \ 30,\ \n "city\": \"New York\ "},\n{\"name\":\"Jane\",\"age\":25,\"city\":\"Los Angeles\"},\n{\"name\":\"Mike\",\"age\":35,\"city\":\"Chicago\"}\]`

	// Remover as barras invertidas antes de colchetes e chaves
	escapedJSONString = strings.ReplaceAll(escapedJSONString, `\n`, "")
	escapedJSONString = strings.ReplaceAll(escapedJSONString, `\`, "")

	// Array de mapas genéricos para armazenar o JSON
	var result []map[string]interface{}

	// Parsear a string JSON
	err := json.Unmarshal([]byte(escapedJSONString), &result)
	if err != nil {
		fmt.Println("Erro ao parsear o JSON:", err)
		return
	}

	// Serializar novamente o JSON formatado
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar o JSON:", err)
		return
	}

	// Criar arquivo físico .json
	fileName := "output.json"
	file, err := os.Create(fileName)
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

	// Upload para o S3 usando LocalStack
	err = uploadToS3(fileName)
	if err != nil {
		fmt.Println("Erro ao fazer upload para o S3:", err)
		return
	}
	fmt.Println("Upload para o S3 realizado com sucesso!")
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
