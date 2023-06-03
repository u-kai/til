package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)




// make main func to run dynamodb get item api to localstack

func main() {
	// create dynamodb client
	// create context
	// create getiteminput
	// call getitem api
	// print response	
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID && region == "us-east-1" {
					return aws.Endpoint{
						PartitionID:   "aws",
						URL:           "http://localhost:4566",
						SigningRegion: "us-east-1",
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			},
		)),
	)
	if err != nil {
		panic(err)
	}
	client := dynamodb.NewFromConfig(cfg)
	ctx := context.TODO()
	input := &dynamodb.GetItemInput{
		TableName: aws.String("test"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: "1",
			},
		},
	}
	resp, err := client.GetItem(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
