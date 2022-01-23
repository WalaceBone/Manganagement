package main

import (
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
	case "DELETE":
		return deleteCharacter(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

// 	resp, err := http.Get(DefaultHTTPGetAddress)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{}, err
// 	}

// 	if resp.StatusCode != 200 {
// 		return events.APIGatewayProxyResponse{}, ErrNon200Response
// 	}

// 	ip, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{}, err
// 	}

// 	if len(ip) == 0 {
// 		return events.APIGatewayProxyResponse{}, ErrNoIP
// 	}

// 	return events.APIGatewayProxyResponse{
// 		Body:       "hello",
// 		StatusCode: 200,
// 	}, nil
// }

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

func deleteCharacter(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	name := req.QueryStringParameters["name"]
	desc := req.QueryStringParameters["desc"]
	_, err := deleteItem(name, desc)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(router)
}
