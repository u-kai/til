package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	v4 "github.com/aws/aws-sdk-go-v2/signer"
)

func main() {
	ctx := context.Background()

	// AWSの設定を読み込む
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// リージョン
	region := cfg.Region

	// リクエストのURLとHTTPメソッド
	apiId := "your-api-id"
	stage := "your-stage"
	resource := "your-resource-path"
	url := fmt.Sprintf("https://%s.execute-api.%s.amazonaws.com/%s/%s", apiId, region, stage, resource)
	method := "GET"

	// リクエストを作成
	req, _ := http.NewRequest(method, url, nil)

	// 署名の作成
	signer := v4.NewSigner()
	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		fmt.Printf("Error retrieving credentials: %v\n", err)
		return
	}
	_, err = signer.SignHTTP(ctx, credentials, req, []byte{}, "execute-api", region, time.Now())

	if err != nil {
		fmt.Printf("Error signing request: %v\n", err)
		return
	}

	// 署名付きリクエストの実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
		return

	}

}
