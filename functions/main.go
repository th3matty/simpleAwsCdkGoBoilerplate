package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jsoniter "github.com/json-iterator/go"

	//"github.com/th3matty/simpleAwsCdkGoBoilerplate/functions/myGoLambda/model/"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	lambda.Start(handle)
}

func init() {
	if os.Getenv("ENVIRONMENT") == "live" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input model.InputPost
	err := jsoniter.UnmarshalFromString(request.Body, &input)
	if err != nil {
		logger.Error("Error unmarshaling body", zap.Error(err))
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	// reult logic here
	// fmt.Println(strings.ContainsAny("team", "i"))
	resultJson, err := jsoniter.MarshalToString(result)

	if err != nil {
		logger.Error("Error marshaling result", zap.Error(err))
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            resultJson,
		IsBase64Encoded: false,
	}, nil
}
