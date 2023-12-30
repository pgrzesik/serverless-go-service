package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
	return "hello there!", nil
}

func main() {
	lambda.Start(handler)
}
