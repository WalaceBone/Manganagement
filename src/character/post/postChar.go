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

func putItem(charac *Character) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("manganagement-live"),
		Item: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(charac.ID),
			},
			"name": {
				S: aws.String(charac.Name),
			},
			"title": {
				S: aws.String(charac.Title),
			},
			"description": {
				S: aws.String(charac.Description),
			},
			"image": {
				S: aws.String(charac.Image),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}
