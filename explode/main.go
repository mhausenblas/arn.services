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
	// NB that the following currently craps out with resources that use a `/`
	// to separate resource type and ID, for example, IAM for user/abc:
	as, ok := request.PathParameters["arnstr"]
	if !ok {
		return serverError(fmt.Errorf("No input provided"))
	}
	if !arn.IsARN(as) {
		return serverError(fmt.Errorf("Need an ARN to operate on"))
	}
	fmt.Printf("Input ARN: %v\n", as)
	a, err := arn.Parse(as)
	if err != nil {
		return serverError(fmt.Errorf("Can't parse %v as an ARN due to: %v", as, err))
	}
	ajson, err := json.Marshal(a)
	if err != nil {
		return serverError(fmt.Errorf("Can't marshal result due to %v", err))
	}
	fmt.Println(string(ajson))
	fmt.Printf("END EXPLODE FUNC\n")
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
