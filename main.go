package main

import (
	"encoding/json"
	"fmt"

	b64 "encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	qrcode "github.com/skip2/go-qrcode"
)

// Response message
type Response struct {
	Message string `json:"msg"`
	Qrcode  string `json:"qr"`
}

// Handler for output
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	// create qr code
	png, err := qrcode.Encode("https://www.smile-feedback.de/vote/1234", qrcode.Medium, 256)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	// create base64 string
	uEnc := b64.URLEncoding.EncodeToString(png)

	r := Response{
		Message: "Hello from golang function",
		Qrcode:  uEnc,
	}
	rbytes, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(rbytes), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
