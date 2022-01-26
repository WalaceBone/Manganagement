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

func getAllItem() ([]*Character, error) {

	//input := &dynamodb.GetItemInput{}
	//if name empty --> search if desc CONTAINS desc
	// if desc empty search by name

	input := &dynamodb.ScanInput{
		TableName: aws.String("manganagement-live"),
	}
	result, err := db.Scan(input)
	if err != nil {
		return nil, err
	}
	if result.Items == nil {
		return nil, nil
	}

	charList := make([]*Character, 1)
	for _, i := range result.Items {
		item := Character{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		charList = append(charList, &item)
	}

	if err != nil {
		return nil, err
	}
	return charList, nil
}
