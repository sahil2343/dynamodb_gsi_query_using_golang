package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	PersonSSN string
	Email     string
	FirstName string
	LastName  string
}

func main() {
	// STarting dynamoDB session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Here TableName is people, partiton key is PersonSSN and index is User_mail
	tableName := "people"
	PersonSSN := "11111111"
	// User_mail := "abcd@gmail.com"
	var resp, err = svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PersonSSN": {
				S: aws.String(PersonSSN),
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	// ************************************************************************************************************************
	// result, err := svc.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String(tableName),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"PersonSSN": {
	// 			S: aws.String(PersonSSN),
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("Got error calling GetItem: %s", err)
	// }

	// if result.Item == nil {
	// 	return
	// }

	// item := Item{}

	// err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	// }
	// fmt.Println("Found item:")
	// fmt.Println("Email:  ", item.Email)
	// fmt.Println("Firstname: ", item.FirstName)
	// fmt.Println("Lastname:  ", item.LastName)

	// proj := expression.NamesList(expression.Name("PersonSSN")
	// keyCondition := expression.Key("PersonSSN").Equal(expression.Value(PersonSSN))
	// expr, errExpression := expression.NewBuilder().WithKeyCondition(keyCondition).WithProjection(proj).Build()

	// params := &dynamodb.QueryInput{
	// 	ExpressionAttributeValues: expr.Values(),
	// 	ExpressionAttributeNames:  expr.Names(),
	// 	ProjectionExpression:      expr.Projection(),
	// 	TableName:                 aws.String(tableName),
	// 	IndexName:                 aws.String("Email-index"),
	// 	KeyConditionExpression:    expr.KeyCondition(),
	// },
	// result, errResults = svc.Query(params)
	// fmt.Print(errResults)
	// fmt.Print(result)

}
