package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) error {
	log.Printf("START lambda function")
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		fmt.Printf("Recieve SNS: [%s/%s] %s\n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		alert := alertTable{
			Timestamp: snsRecord.Timestamp,
			Event:     record.EventSource,
			Message:   snsRecord.Message,
		}
		if err := putDynamo(alert); err != nil {
			return fmt.Errorf("dynamoDB put data error: %v", err.Error())
		}
		fmt.Print("insert dynamoDB item")
	}
	log.Printf("END lambda function")
	return nil
}

func main() {
	lambda.Start(handler)
}
