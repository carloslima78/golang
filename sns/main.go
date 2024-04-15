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
