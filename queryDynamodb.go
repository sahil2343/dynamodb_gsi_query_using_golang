package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	Phone     string
}

func queryHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// STarting dynamoDB session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Here TableName is people, partiton key is PersonSSN and index is User_mail
	tableName := "people"
	// PersonSSN := "11111111"
	phone_no := "12347"
	index_name := "Phone-index"

	//query using GSI(indexing)
	var queryInput, err2 = svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String(index_name),
		KeyConditions: map[string]*dynamodb.Condition{
			"Phone": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(phone_no),
					},
				},
			},
		},
	})

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(queryInput)
	var item Item
	dynamodbattribute.UnmarshalMap(queryInput.Items[0], &item)
	jsonData, _ := json.Marshal(item)
	fmt.Println(string(jsonData))

	return events.APIGatewayProxyResponse{
		Body:       string(jsonData),
		StatusCode: 200,
	}, nil

	// ************************************************************************************************************************
	// query using partition key
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

}

func main() {
	lambda.Start(queryHandler)
}
