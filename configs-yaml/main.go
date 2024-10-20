package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Struct para receber as configurações mapeadas do arquivo YAML
type Config struct {
	HTTPPort     string `yaml:"http_port"`
	DBEngineer   string `yaml:"db_engineer"`
	DBConnection string `yaml:"db_connection"`
	TableName    string `yaml:"table_name"`
}

func main() {
	// Leitura do arquivo YAML
	data, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo YAML: %v", err)
	}

	// Desserialização do arquivo YAML via Unmarshal para a struct Config
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do YAML: %v", err)
	}

	// Exibição dos valores
	fmt.Printf("Porta HTTP: %s\n", config.HTTPPort)
	fmt.Printf("Banco de Dados: %s\n", config.DBEngineer)
	fmt.Printf("Conexão: %s\n", config.DBConnection)
	fmt.Printf("Nome da tabela: %s\n", config.TableName)
}
