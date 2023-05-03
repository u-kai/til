package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

func AddAwsSignature(c *http.Client, req *http.Request) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", err
	}
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	fmt.Printf("creds : %v\n", creds.SecretAccessKey)
	if err != nil {
		return "", err
	}
	signer := v4.NewSigner()
	payloadHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	err = signer.SignHTTP(context.TODO(), creds, req, payloadHash, "apigateway", *&cfg.Region, time.Now().UTC())

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
