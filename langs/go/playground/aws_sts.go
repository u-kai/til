package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	conf := aws.NewConfig()
	sess, _ := session.NewSession(conf)
	client := sts.New(sess)
	result, _ := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	fmt.Printf("%s", result)
}
