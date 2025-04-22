package main

import (
	"context"
	"github.com/aidenismee/go-lambda-sample/internal/server"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambdaHandler *echoadapter.EchoLambda

func init() {
	echoLambdaHandler = server.InitLambdaFunctionHandler()
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambdaHandler.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
