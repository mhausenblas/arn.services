package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/arn"
)

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Println(err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: fmt.Sprintf("%v", err.Error()),
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("START EXPLODE FUNC\n")
	as := request.Path
	a, err := arn.Parse(as)
	if err != nil {
		return serverError(err)
	}
	fmt.Printf("END EXPLODE FUNC\n")
	ajson, err := json.Marshal(a)
	if err != nil {
		return serverError(fmt.Errorf("Can't marshal result: %v", err))
	}
	fmt.Println(string(ajson))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		Body: string(ajson),
	}, nil
}

func main() {
	lambda.Start(handler)
}
