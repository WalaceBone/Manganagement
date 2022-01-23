package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Character struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var db = dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("eu-west-3"))

func deleteItem(name string, desc string) (*Character, error) {

	//input := &dynamodb.GetItemInput{}
	//if name empty --> search if desc CONTAINS desc
	// if desc empty search by name

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("manganagement-live"),
		Key: map[string]*dynamodb.AttributeValue{
			"description": {
				S: aws.String(desc),
			},
			"name": {
				S: aws.String(name),
			},
		},
	}
	_, err := db.DeleteItem(input)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
