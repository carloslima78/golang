package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// Criando uma nova sessão da AWS
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Criando um novo cliente SQS
	sqsClient := sqs.New(sess)

	// URL da fila de origem (myqueue-1)
	sourceQueueURL := "https://sqs.us-east-1.amazonaws.com/303315406913/myqueue1"

	// URL da fila de destino (myqueue2)
	destinationQueueURL := "https://sqs.us-east-1.amazonaws.com/303315406913/myqueue2"

	// Canal para transmitir mensagens da fila de origem para a fila de destino
	messageChannel := make(chan string)

	// Iniciar a go routine para consumir mensagens e enviar para o canal
	go consumer(sqsClient, sourceQueueURL, messageChannel)

	// Iniciar a go routine para receber mensagens do canal e produzir na fila de destino
	go producer(sqsClient, destinationQueueURL, messageChannel)

	// Mantendo o programa em execução
	select {}
}

func consumer(svc *sqs.SQS, sourceQueueURL string, messageChannel chan<- string) {
	for {
		// Recebendo uma única mensagem da fila de origem
		receiveParams := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sourceQueueURL),
			MaxNumberOfMessages: aws.Int64(1), // Receber apenas uma mensagem
			WaitTimeSeconds:     aws.Int64(20),
		}

		receiveResp, err := svc.ReceiveMessage(receiveParams)
		if err != nil {
			log.Printf("Erro ao receber mensagem: %v", err)
			continue
		}

		if len(receiveResp.Messages) == 0 {
			log.Println("Não há mensagens na fila de origem.")
			continue
		}

		// Lendo a mensagem recebida
		receivedMessage := *receiveResp.Messages[0].Body
		fmt.Println("Mensagem Recebida:", receivedMessage)

		// Deletando a mensagem da fila de origem
		deleteParams := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(sourceQueueURL),
			ReceiptHandle: receiveResp.Messages[0].ReceiptHandle,
		}

		_, err = svc.DeleteMessage(deleteParams)
		if err != nil {
			log.Printf("Erro ao deletar mensagem: %v", err)
			continue
		}

		fmt.Println("Mensagem recebida e deletada com sucesso:", receivedMessage)

		// Enviar mensagem para o canal
		messageChannel <- receivedMessage
	}
}

func producer(svc *sqs.SQS, destinationQueueURL string, messageChannel <-chan string) {
	for {
		// Receber mensagem do canal
		message, ok := <-messageChannel
		if !ok {
			log.Println("Canal fechado. Encerrando a go routine do produtor.")
			return
		}

		// Enviando a mensagem para a fila de destino
		sendParams := &sqs.SendMessageInput{
			MessageBody:  aws.String(message),
			QueueUrl:     aws.String(destinationQueueURL),
			DelaySeconds: aws.Int64(0),
		}

		_, err := svc.SendMessage(sendParams)
		if err != nil {
			log.Printf("Erro ao enviar mensagem para a fila de destino: %v", err)
			continue
		}

		fmt.Println("Mensagem enviada para a fila de destino com sucesso:", message)

		// Aguardando um segundo antes de enviar a próxima mensagem
		// Isso é apenas um exemplo, ajuste conforme necessário
		time.Sleep(time.Second)
	}
}
