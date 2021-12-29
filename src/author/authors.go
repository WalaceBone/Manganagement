package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var db = dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("eu-west-3"))

func getItem(name string) (*Author, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("author"),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
		},
	}
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	auth := new(Author)
	err = dynamodbattribute.UnmarshalMap(result.Item, auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func putItem(auth *Author) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("author"),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(auth.ID),
			},
			"Name": {
				S: aws.String(auth.Name),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}
