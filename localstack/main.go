package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func makeDynamoDBSvc () *dynamodb.Client{
	resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				URL: "http://localhost:4566",
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg,err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithEndpointResolver(resolver),
	)
	if err !=nil{
		panic(err)
	}	
	return  dynamodb.NewFromConfig(cfg)
}

func main() {
	svc := makeDynamoDBSvc()
	result,err := svc.ListTables(context.Background(),&dynamodb.ListTablesInput{})
	if err != nil {
		panic(err)
	}
	for _,table := range result.TableNames{
		fmt.Println(table)
	}	
	svc.PutItem(context.Background(),&dynamodb.PutItemInput{
		TableName: aws.String("test"),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
			"name": &types.AttributeValueMemberS{Value: "test"},
		},
	})

	getItem,err := svc.GetItem(context.Background(),&dynamodb.GetItemInput{
		TableName: aws.String("test"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
	})
	if err != nil {
		panic(err)
	}
	for k,v := range getItem.Item{
		fmt.Println(k,v)
	}
}

