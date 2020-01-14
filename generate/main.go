package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/arn"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	fmt.Printf("START GENERATE FUNC\n")
	a := arn.ARN{}
	err := json.Unmarshal([]byte(request.Body), &a)
	if err != nil {
		return serverError(fmt.Errorf("Can't parse %v as an ARN due to: %v", request.Body, err))
	}
	// a somewhat sensible defaulting for fields:
	switch {
	case a.Partition == "":
		a.Partition = "aws"
	// NB this has to be fixed to take non-regional services such as IAM into
	// account, but for now let's keep it simple:
	case a.Region == "":
		a.Region = "us-west-2"
	}
	fmt.Printf("END GENERATE FUNC\n")
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "text/plain",
		},
		Body: a.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
