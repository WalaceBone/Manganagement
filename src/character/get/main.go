package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("no ip in http response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("non 200 response found")
)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return showCharacter(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func showCharacter(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	name := req.QueryStringParameters["name"]
	desc := req.QueryStringParameters["desc"]
	charact, err := getItem(name, desc)
	if err != nil {
		return serverError(err)
	}

	jres, err := json.Marshal(charact)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jres),
	}, nil
}

func main() {
	lambda.Start(router)
}
