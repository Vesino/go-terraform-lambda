package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Profile struct {
	Email string `json:"email"`
}

func NewProfile(email string) *Profile {
	return &Profile{
		Email: email,
	}
}

// DB represents a structure to control DynamoDB interactions
type DB struct {
	cfg    aws.Config
	client *dynamodb.Client
}

// NewDB creates a new instance of *DB based on provided options
func NewDB(options ...func(*config.LoadOptions) error) (*DB, error) {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		options...,
	)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DB{
		client: client,
	}, nil
}

// CreateProfile persist the profile into dynamodb instance
func (db *DB) CreateProfile(ctx context.Context, profile *Profile) error {
	// transform profile object into a dynamodb compatible item
	item, err := attributevalue.MarshalMap(profile)
	if err != nil {
		return err
	}

	_, err = db.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("tutorial_profiles"),
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}
