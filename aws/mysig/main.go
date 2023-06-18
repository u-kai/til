package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

// q:aws のv4の署名を作成するにはどうすればいいのか？
// a: https://docs.aws.amazon.com/ja_jp/general/latest/gr/sigv4-signed-request-examples.html
func main() {

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile("default"),
	)
	if err != nil {
		panic(err.Error())
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		panic(err.Error())
	}

	type Test struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
	data := Test{Id: "1", Name: "test"}
	b, err := json.Marshal(data)
	result := sha256.Sum256(b)
	hash := hex.EncodeToString(result[:])
	req, err := http.NewRequest(
		http.MethodPost,
		"https://x.execute-api.ap-northeast-1.amazonaws.com/prod",
		bytes.NewReader(b),
	)
	if err != nil {
		panic(err.Error())
	}

	hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	signer := v4.NewSigner()
	signer.SignHTTP(ctx, credentials, req, hash, "execute-api", "ap-northeast-1", time.Now())

	if err != nil {
		panic(err.Error())
	}

	httpClient := new(http.Client)

	response, err := httpClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(string(responseBody))
}
