package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Manga struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

var db = dynamodb.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion("eu-west-3"))

func getItem(name string) (*Manga, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("mangater"),
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
	manga := new(Manga)
	err = dynamodbattribute.UnmarshalMap(result.Item, manga)
	if err != nil {
		return nil, err
	}
	return manga, nil
}

func putItem(manga *Manga) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("mangater"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(manga.ID),
			},
			"name": {
				S: aws.String(manga.Title),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}
