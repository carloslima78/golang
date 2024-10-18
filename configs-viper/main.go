package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Defina sua struct de configuração
type Config struct {
	DBConnection string `yaml:"db_connection"`
	TableName    string `yaml:"table_name"`
}

func main() {
	// Leia o arquivo YAML
	data, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo YAML: %v", err)
	}

	// Unmarshal o YAML para a struct Config
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Erro ao fazer unmarshal do YAML: %v", err)
	}

	// Exiba os valores
	fmt.Printf("Conexão: %s\n", config.DBConnection)
	fmt.Printf("Nome da tabela: %s\n", config.TableName)
}
