package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Character struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var db = dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("eu-west-3"))

func getItem(name string, desc string) (*Character, error) {

	//input := &dynamodb.GetItemInput{}
	//if name empty --> search if desc CONTAINS desc
	// if desc empty search by name

	input := &dynamodb.GetItemInput{
		TableName: aws.String("manganagement-live"),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
			"description": {
				S: aws.String(desc),
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
