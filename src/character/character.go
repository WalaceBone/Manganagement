package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Character struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Manga       string `json:"manga"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var db = dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("eu-west-3"))

func getItem(name string) (*Character, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("character"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
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
	charac := new(Character)
	err = dynamodbattribute.UnmarshalMap(result.Item, charac)
	if err != nil {
		return nil, err
	}
	return charac, nil
}

func putItem(charac *Character) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("character"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(charac.ID),
			},
			"name": {
				S: aws.String(charac.Name),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}
