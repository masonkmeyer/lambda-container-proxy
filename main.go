package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	proxy "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *proxy.EchoLambda

func init() {
	e := echo.New()

	e.GET("/ints", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []int{1, 2, 3})
	})

	e.GET("/strings", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, []string{"A", "B", "C"})
	})

	echoLambda = proxy.New(e)
}

// handler is the lambda handler for an APIProxyRequest
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
