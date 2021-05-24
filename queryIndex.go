package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	PersonSSN string
	Email     string
	FirstName string
	LastName  string
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "people"
	PersonSSN := "11111111"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PersonSSN": {
				S: aws.String(PersonSSN),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		return
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	fmt.Println("Found item:")
	fmt.Println("Email:  ", item.Email)
	fmt.Println("Firstname: ", item.FirstName)
	fmt.Println("Lastname:  ", item.LastName)

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		Email *string `min:"3" type:"string"`

	}
}
