package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Input struct {
	Email string `json:"email"`
}

func HandleRequest(ctx context.Context, event *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	var err error

	body := []byte(event.Body)
	if event.IsBase64Encoded {
		body, err = base64.StdEncoding.DecodeString(string(body))
		if err != nil {
			return &events.APIGatewayV2HTTPResponse{
				StatusCode: 400,
				Body:       "Cannot process the request",
				Headers: map[string]string{
					"Content-Type": "text/plain",
				},
			}, err
		}
	}

	var input *Input
	if err := json.Unmarshal(body, &input); err != nil {
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Cannot process the request",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	}

	db, err := NewDB()
	if err != nil {
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Cannot process the request at the moment",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	}

	profile := NewProfile(input.Email)
	if err := db.CreateProfile(ctx, profile); err != nil {
		fmt.Println("Error creating profile", err)

		return &events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Cannot process the request at the moment",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 204,
		Body:       "",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil

}
