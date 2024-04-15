# SNS

## Localstack


### Baixando a imagem Docker

``` bash
docker pull localstack/localstack
```


### Executando o container Docker

``` bash
docker run -it -d -p 4566:4566 --name localstack localstack/localstack
```


### Criando um tópico SNS 

Inicie o comando abaixo via aws cli

``` bash
aws --endpoint-url=http://localhost:4566 sns create-topic --name meu-topico
```


### Criando um módulo para aplicação

``` bash
go mod init sns
```

### Instalando a dependência do AWS SDK

Para instalar a dependência do SDK da AWS em um projeto Go, você pode usar o go get. Aqui está o comando básico para instalar o SDK da AWS:

``` bash
go get github.com/aws/aws-sdk-go
```

## Criando a aplicação

``` go
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	svc := sns.New(sess)

	publishParams := &sns.PublishInput{
		Message:  aws.String("Olá mundo bom"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:meu-topico"),
	}

	_, err := svc.Publish(publishParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mensagem enviada")
}
```

- Execute o comando abaixo para iniciar a aplicação

``` bash
go run main.go
```

## Parando e removendo o container localstack

``` bash
docker stop localstack

docker rm localstack
```