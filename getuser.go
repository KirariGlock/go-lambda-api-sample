package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type User struct {
	UserID   string `json:"userid"`
	USerName string `json:"username"`
}

type Request struct {
	ID string `json:"id"`
}

type Response struct {
	Data string `json:"body"`
}

func Handler(request Request) (Response, error) {

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("User"),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(request.ID),
			},
		},
	}

	getItem, getErr := svc.GetItem(getParams)
	if getErr != nil {
		panic(getErr)
	}

	return Response{
		Data: fmt.Sprintln(getItem),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
