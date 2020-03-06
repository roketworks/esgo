package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/roketworks/serverless-es-go/package"
)

var awsSession = session.Must(session.NewSession(&aws.Config{}))
var snsSvc = sns.New(awsSession)
var dynamoSvc = dynamodb.New(awsSession)
var es = &_package.DynamoDbEventStore{Db: dynamoSvc, EventTable: "es-test-stack-EventStoreTable-ONSDU4YWZ62N"}

func handler(e events.SNSEvent) error {
	for _, record := range e.Records {
		// TODO: map
		_, err := _package.Save(es, "", 1, []byte(record.SNS.Message))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
